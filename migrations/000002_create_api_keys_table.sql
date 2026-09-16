-- +goose Up
CREATE TABLE api_keys (
    id BIGSERIAL PRIMARY KEY,
    owner_email TEXT NOT NULL UNIQUE,
    key_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    last_used_at TIMESTAMP NULL,
    revoked_at TIMESTAMP NULL
);

-- +goose Down
DROP TABLE IF EXISTS api_keys CASCADE;
