package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Two token types, because they answer different questions.
//
// The access token answers "who is this request from?" on every single call,
// so it must be verifiable without touching the database. It is a signed JWT
// with a short life, and it is not stored anywhere.
//
// The refresh token answers "may this session continue?" occasionally, so it
// can afford a database read, and because it lives for weeks it has to be
// revocable. It is an opaque random string, stored only as a hash.
//
// A single long-lived token would have to be either unrevocable or checked
// against the database on every request. This split avoids both.
const (
	AccessTokenTTL  = 15 * time.Minute
	RefreshTokenTTL = 30 * 24 * time.Hour
)

var (
	ErrInvalidToken = errors.New("token is not valid")
	ErrExpiredToken = errors.New("token has expired")
)

// Claims is what an access token carries. Deliberately minimal: an identifier
// and the standard registered claims. Anything else, display name, role,
// membership, would be a copy of database state that goes stale the moment it
// is issued, and callers would trust the stale copy.
type Claims struct {
	jwt.RegisteredClaims
}

// Issuer mints and verifies access tokens.
type Issuer struct {
	secret []byte
	now    func() time.Time // injected so expiry is testable without sleeping
}

// NewIssuer fails on a short secret rather than warning about it.
//
// A 32-byte minimum for HS256 is the RFC 7518 requirement: the key should be
// at least as long as the hash output, or the signature is weaker than the
// algorithm implies. A deployment that starts with a weak signing key would
// look healthy while being trivially forgeable, so it does not start.
func NewIssuer(secret string) (*Issuer, error) {
	if len(secret) < 32 {
		return nil, fmt.Errorf("auth: signing secret must be at least 32 bytes, got %d", len(secret))
	}
	return &Issuer{secret: []byte(secret), now: time.Now}, nil
}

// Issue returns a signed access token for a user.
func (i *Issuer) Issue(userID uuid.UUID) (string, time.Time, error) {
	now := i.now()
	expires := now.Add(AccessTokenTTL)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expires),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "skriptra",
			ID:        uuid.NewString(),
		},
	})

	signed, err := token.SignedString(i.secret)
	if err != nil {
		return "", time.Time{}, err
	}
	return signed, expires, nil
}

// Verify checks a token and returns the user it identifies.
//
// The signing method is pinned to HMAC. Without that check a token whose
// header says `"alg": "none"`, or one signed with the public half of an RSA
// pair, would be accepted: the classic algorithm-confusion attack, and the
// reason this is not a one-line parse call.
func (i *Issuer) Verify(tokenString string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
			}
			return i.secret, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer("skriptra"),
		jwt.WithTimeFunc(i.now),
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return uuid.Nil, ErrExpiredToken
		}
		return uuid.Nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return uuid.Nil, ErrInvalidToken
	}

	id, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, ErrInvalidToken
	}
	return id, nil
}

// NewRefreshToken returns a fresh opaque token and the hash to store for it.
//
// 32 bytes from crypto/rand, so it is unguessable, and no structure at all, so
// there is nothing in it to tamper with. The server keeps only the hash, which
// means a database dump does not hand an attacker a working session.
//
// SHA-256 rather than argon2 here: the input is 256 bits of entropy already,
// so there is no dictionary to attack and nothing for a slow hash to buy.
func NewRefreshToken() (token, hash string, err error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", "", fmt.Errorf("read random: %w", err)
	}
	token = base64.RawURLEncoding.EncodeToString(buf)
	return token, HashRefreshToken(token), nil
}

// HashRefreshToken is the lookup key for a stored session.
func HashRefreshToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
