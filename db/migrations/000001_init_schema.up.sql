CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "line_id" varchar NOT NULL UNIQUE,
  "email" varchar NOT NULL UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("username");
