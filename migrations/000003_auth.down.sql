BEGIN;

DROP TABLE IF EXISTS refresh_tokens;

ALTER TABLE users DROP CONSTRAINT IF EXISTS users_email_key;
ALTER TABLE users ALTER COLUMN email DROP NOT NULL;
ALTER TABLE users DROP COLUMN IF EXISTS password_hash;

-- The placeholder addresses minted by the up migration are removed again, so
-- rolling forward and back twice does not leave rows claiming to own an email
-- nobody chose.
UPDATE users SET email = NULL WHERE email LIKE '%@placeholder.invalid';

COMMIT;
