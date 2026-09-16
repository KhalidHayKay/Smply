-- +goose Up
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    original TEXT NOT NULL,
    alias TEXT UNIQUE NOT NULL,
    visited INTEGER DEFAULT 0,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_visited TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS urls CASCADE;
