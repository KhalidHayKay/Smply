-- +goose Up
ALTER TABLE api_keys DROP CONSTRAINT api_keys_owner_email_key;

-- +goose Down
ALTER TABLE api_keys ADD CONSTRAINT api_keys_owner_email_key UNIQUE (owner_email);
