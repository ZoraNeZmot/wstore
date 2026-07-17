CREATE TABLE IF NOT EXISTS accounts.tokens (
    token text PRIMARY KEY,
    user_uid uuid NOT NULL REFERENCES accounts.users(uid) ON DELETE CASCADE,
    valid boolean NOT NULL DEFAULT TRUE,
    expires_at timestamp NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);