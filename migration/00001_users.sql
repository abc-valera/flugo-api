-- +goose Up
CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "fullname" varchar NOT NULL DEFAULT '',
  "status" varchar NOT NULL DEFAULT '',
  "bio" varchar NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

-- +goose Down
DROP TABLE IF EXISTS "users";