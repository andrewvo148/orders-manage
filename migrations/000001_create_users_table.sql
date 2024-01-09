-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id         text        NOT NULL,
  full_name       text        NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);


-- +goose Down
DROP TABLE IF EXISTS users;
