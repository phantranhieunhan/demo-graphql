CREATE TABLE IF NOT EXISTS "links"(
  "id" bigserial PRIMARY KEY,
    "title" varchar,
    "address" varchar,
    "user_id" bigserial NOT NULL
);

ALTER TABLE "links" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
