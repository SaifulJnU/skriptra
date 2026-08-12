BEGIN;

-- Real authentication.
--
-- The schema always had a users table and a course_members table with roles,
-- because authorization checks were wired from the first endpoint and only the
-- identity behind them was a stub. This migration fills in the identity: a
-- credential to prove who you are, and a place to keep sessions.

-- A password hash, not a password. Nullable because the seeded development
-- user predates this and has no credential; it simply cannot log in.
ALTER TABLE users ADD COLUMN IF NOT EXISTS password_hash text;

-- Email becomes the login handle, so it has to identify exactly one account.
-- The column is already citext, so uniqueness is case insensitive and
-- Saiful@example.com cannot register alongside saiful@example.com.
UPDATE users SET email = id::text || '@placeholder.invalid' WHERE email IS NULL;

ALTER TABLE users ALTER COLUMN email SET NOT NULL;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'users_email_key'
    ) THEN
        ALTER TABLE users ADD CONSTRAINT users_email_key UNIQUE (email);
    END IF;
END $$;

-- Sessions.
--
-- Access tokens are short-lived JWTs and are deliberately not stored: verifying
-- one is a signature check, and putting a database read in front of every
-- request would undo the reason for using them. Refresh tokens are the opposite
-- case. They live for weeks, so they must be revocable, which means they have
-- to be stored.
--
-- Only the hash is kept. A stolen database dump then yields no usable token,
-- for the same reason passwords are not stored in the clear.
CREATE TABLE refresh_tokens (
    id          uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id     uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash  text NOT NULL UNIQUE,
    expires_at  timestamptz NOT NULL,
    revoked_at  timestamptz,
    created_at  timestamptz NOT NULL DEFAULT now(),
    user_agent  text
);

-- Every lookup is by hash, and every logout-everywhere is by user.
CREATE INDEX refresh_tokens_user_idx ON refresh_tokens (user_id) WHERE revoked_at IS NULL;

COMMIT;
