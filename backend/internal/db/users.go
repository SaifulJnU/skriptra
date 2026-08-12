package db

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

// ErrEmailTaken is returned instead of a raw constraint violation so the
// handler can answer without importing pgx or matching on error text.
var ErrEmailTaken = errors.New("email already registered")

// Account is a user with their credential. Kept out of the domain package on
// purpose: domain types are wire types, and a password hash must never be one
// field rename away from being serialised to a client.
type Account struct {
	ID           uuid.UUID
	Email        string
	DisplayName  string
	PasswordHash string
	CreatedAt    time.Time
}

// CreateAccount registers a new user.
//
// The unique constraint on email is what actually prevents duplicates, not a
// prior SELECT. Two simultaneous signups for the same address both pass a
// check-then-insert and one of them still has to lose at the constraint, so
// the constraint is treated as the decision rather than as an unexpected
// error.
func (s *Store) CreateAccount(ctx context.Context, email, displayName, passwordHash string) (*Account, error) {
	a := Account{
		Email:       strings.TrimSpace(email),
		DisplayName: strings.TrimSpace(displayName),
	}

	err := s.pool.QueryRow(ctx, `
		INSERT INTO users (email, display_name, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`,
		a.Email, a.DisplayName, passwordHash).Scan(&a.ID, &a.CreatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, ErrEmailTaken
		}
		return nil, err
	}
	a.PasswordHash = passwordHash
	return &a, nil
}

// AccountByEmail looks up a login. The email column is citext, so the match is
// case insensitive without lower() defeating the index.
func (s *Store) AccountByEmail(ctx context.Context, email string) (*Account, error) {
	var a Account
	var hash *string
	err := s.pool.QueryRow(ctx, `
		SELECT id, email, display_name, password_hash, created_at
		FROM users WHERE email = $1`, strings.TrimSpace(email)).Scan(
		&a.ID, &a.Email, &a.DisplayName, &hash, &a.CreatedAt)
	if err != nil {
		return nil, norm(err)
	}
	if hash != nil {
		a.PasswordHash = *hash
	}
	return &a, nil
}

// AccountByID resolves the subject of an access token to a live user.
func (s *Store) AccountByID(ctx context.Context, id uuid.UUID) (*Account, error) {
	var a Account
	var hash *string
	err := s.pool.QueryRow(ctx, `
		SELECT id, email, display_name, password_hash, created_at
		FROM users WHERE id = $1`, id).Scan(
		&a.ID, &a.Email, &a.DisplayName, &hash, &a.CreatedAt)
	if err != nil {
		return nil, norm(err)
	}
	if hash != nil {
		a.PasswordHash = *hash
	}
	return &a, nil
}

// ---------------------------------------------------------------- sessions --

// StoreRefreshToken records an issued session.
func (s *Store) StoreRefreshToken(ctx context.Context, userID uuid.UUID, tokenHash string, expiresAt time.Time, userAgent string) error {
	_, err := s.pool.Exec(ctx, `
		INSERT INTO refresh_tokens (user_id, token_hash, expires_at, user_agent)
		VALUES ($1, $2, $3, $4)`, userID, tokenHash, expiresAt, truncate(userAgent, 500))
	return err
}

// RotateRefreshToken exchanges one refresh token for another, atomically.
//
// Rotation on every use is what makes a stolen refresh token survivable: the
// legitimate client refreshes, the stolen copy is now revoked, and the next
// attempt to use it fails. Doing the revoke and the insert in one transaction
// means a crash between them cannot leave a session that is neither valid nor
// replaced.
//
// An already-revoked or expired token returns ErrNotFound. The caller must not
// distinguish those cases to the client, since "this token existed but was
// used" is information an attacker holding a stolen token would find useful.
func (s *Store) RotateRefreshToken(
	ctx context.Context, oldHash, newHash string, expiresAt time.Time, userAgent string,
) (uuid.UUID, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)

	// The row is locked before it is inspected, so two concurrent refreshes
	// with the same token cannot both succeed.
	var userID uuid.UUID
	err = tx.QueryRow(ctx, `
		SELECT user_id FROM refresh_tokens
		WHERE token_hash = $1 AND revoked_at IS NULL AND expires_at > now()
		FOR UPDATE`, oldHash).Scan(&userID)
	if err != nil {
		return uuid.Nil, norm(err)
	}

	if _, err := tx.Exec(ctx,
		`UPDATE refresh_tokens SET revoked_at = now() WHERE token_hash = $1`, oldHash); err != nil {
		return uuid.Nil, err
	}

	if _, err := tx.Exec(ctx, `
		INSERT INTO refresh_tokens (user_id, token_hash, expires_at, user_agent)
		VALUES ($1, $2, $3, $4)`, userID, newHash, expiresAt, truncate(userAgent, 500)); err != nil {
		return uuid.Nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}

// RevokeRefreshToken ends one session. Logging out an already-revoked token is
// not an error: the caller asked for it to be unusable and it is.
func (s *Store) RevokeRefreshToken(ctx context.Context, tokenHash string) error {
	_, err := s.pool.Exec(ctx,
		`UPDATE refresh_tokens SET revoked_at = now()
		 WHERE token_hash = $1 AND revoked_at IS NULL`, tokenHash)
	return err
}

// RevokeAllRefreshTokens ends every session for a user, which is what a
// password change has to do.
func (s *Store) RevokeAllRefreshTokens(ctx context.Context, userID uuid.UUID) error {
	_, err := s.pool.Exec(ctx,
		`UPDATE refresh_tokens SET revoked_at = now()
		 WHERE user_id = $1 AND revoked_at IS NULL`, userID)
	return err
}

// PurgeExpiredRefreshTokens keeps the table from growing without bound. Rows
// are kept for a grace period after expiry so a support question about a
// session can still be answered.
func (s *Store) PurgeExpiredRefreshTokens(ctx context.Context) (int64, error) {
	tag, err := s.pool.Exec(ctx,
		`DELETE FROM refresh_tokens WHERE expires_at < now() - interval '30 days'`)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max]
}
