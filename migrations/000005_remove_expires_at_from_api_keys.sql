-- +goose Up
ALTER TABLE api_keys DROP COLUMN expires_at;

-- +goose Down
ALTER TABLE api_keys ADD COLUMN expires_at TIMESTAMP NOT NULL;
