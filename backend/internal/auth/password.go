// Package auth holds credential handling and token issuance.
//
// Kept separate from the HTTP layer so the security-sensitive parts are in one
// small, readable, testable place rather than spread through handlers.
package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2id, not bcrypt.
//
// bcrypt would be defensible, but it caps the password at 72 bytes and its
// only cost knob is time. Argon2id is the current password-hashing
// recommendation (it won the Password Hashing Competition and is what OWASP
// names first) because it is memory-hard: an attacker with GPUs has to buy
// memory bandwidth, not just cores, which is the expensive thing to scale.
//
// These parameters follow the OWASP minimum of 19 MiB with two iterations.
// They are encoded into every stored hash, so raising them later does not
// invalidate existing passwords: an old hash still verifies under its own
// parameters and can be re-hashed on the next successful login.
const (
	argonTime    uint32 = 2
	argonMemory  uint32 = 19 * 1024 // KiB
	argonThreads uint8  = 1
	argonKeyLen  uint32 = 32
	saltLen             = 16
)

var (
	ErrInvalidHash        = errors.New("password hash is malformed")
	ErrIncompatibleFormat = errors.New("password hash uses an unsupported algorithm")
)

// HashPassword returns a PHC-format string carrying the algorithm, its
// parameters, the salt and the digest.
//
// Self-describing on purpose: the parameters live with the hash rather than in
// a constant that a future change would silently break every stored password
// against.
func HashPassword(password string) (string, error) {
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("read salt: %w", err)
	}

	digest := argon2.IDKey([]byte(password), salt, argonTime, argonMemory, argonThreads, argonKeyLen)

	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, argonMemory, argonTime, argonThreads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(digest),
	), nil
}

// VerifyPassword reports whether a password matches a stored hash.
//
// The comparison is constant time. A byte-by-byte compare that returns early
// leaks how much of the digest matched through timing, which over enough
// samples is enough to reconstruct it.
func VerifyPassword(password, encoded string) (bool, error) {
	params, salt, digest, err := decodeHash(encoded)
	if err != nil {
		return false, err
	}

	candidate := argon2.IDKey([]byte(password), salt,
		params.time, params.memory, params.threads, uint32(len(digest)))

	return subtle.ConstantTimeCompare(digest, candidate) == 1, nil
}

type argonParams struct {
	memory  uint32
	time    uint32
	threads uint8
}

func decodeHash(encoded string) (argonParams, []byte, []byte, error) {
	var p argonParams

	parts := strings.Split(encoded, "$")
	if len(parts) != 6 {
		return p, nil, nil, ErrInvalidHash
	}
	if parts[1] != "argon2id" {
		return p, nil, nil, ErrIncompatibleFormat
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return p, nil, nil, ErrInvalidHash
	}
	if version != argon2.Version {
		return p, nil, nil, ErrIncompatibleFormat
	}

	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &p.memory, &p.time, &p.threads); err != nil {
		return p, nil, nil, ErrInvalidHash
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(parts[4])
	if err != nil {
		return p, nil, nil, ErrInvalidHash
	}
	digest, err := base64.RawStdEncoding.Strict().DecodeString(parts[5])
	if err != nil {
		return p, nil, nil, ErrInvalidHash
	}

	return p, salt, digest, nil
}

// dummyHash is verified against when no account exists for the submitted
// email.
//
// Returning early on an unknown address makes the failed-login response
// measurably faster than one for a real account, which turns the login form
// into an oracle for which addresses are registered. Hashing anyway costs the
// same time and gives the same answer to both.
var dummyHash string

func init() {
	h, err := HashPassword("skriptra-timing-equaliser")
	if err != nil {
		panic("auth: cannot hash the dummy password: " + err.Error())
	}
	dummyHash = h
}

// EqualiseTiming burns the same work a real verification would, for use when
// the account does not exist.
func EqualiseTiming(password string) {
	_, _ = VerifyPassword(password, dummyHash)
}
