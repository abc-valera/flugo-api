-- +goose Up
CREATE TABLE "jokes" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "title" varchar NOT NULL,
  "text" varchar NOT NULL,
  "explanation" varchar NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "jokes" ("username", "title");

ALTER TABLE "jokes" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

-- +goose Down
DROP TABLE IF EXISTS "jokes";
