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

CREATE TABLE "jokes" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "title" varchar NOT NULL,
  "text" varchar NOT NULL,
  "explanation" varchar NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "likes" (
  "username" varchar,
  "joke_id" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("username", "joke_id")
);

CREATE TABLE "comments" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "joke_id" int NOT NULL,
  "text" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "jokes" ("username", "title");

ALTER TABLE "jokes" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "likes" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "likes" ADD FOREIGN KEY ("joke_id") REFERENCES "jokes" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "comments" ADD FOREIGN KEY ("joke_id") REFERENCES "jokes" ("id");