-- +goose Up
CREATE SCHEMA IF NOT EXISTS bedrok;

-- +goose Down
DROP SCHEMA IF EXISTS bedrok CASCADE;
