-- +goose Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT encode(digest(gen_random_uuid()::text, 'sha256'), 'hex');

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;
DROP EXTENSION IF EXISTS pgcrypto;