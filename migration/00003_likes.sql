-- +goose Up
CREATE TABLE "likes" (
  "username" varchar,
  "joke_id" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("username", "joke_id")
);

ALTER TABLE "likes" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "likes" ADD FOREIGN KEY ("joke_id") REFERENCES "jokes" ("id");

-- +goose Down
DROP TABLE IF EXISTS "likes";
