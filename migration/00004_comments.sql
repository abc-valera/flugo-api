-- +goose Up
CREATE TABLE "comments" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "joke_id" int NOT NULL,
  "text" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "comments" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "comments" ADD FOREIGN KEY ("joke_id") REFERENCES "jokes" ("id");

-- +goose Down
DROP TABLE IF EXISTS "comments";
