CREATE TABLE IF NOT EXISTS accounts.tokens (
    token text PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES accounts.users(id) ON DELETE CASCADE,
    is_valid boolean NOT NULL DEFAULT TRUE,
    expires_at timestamp NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);