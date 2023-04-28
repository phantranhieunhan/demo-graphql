CREATE TABLE IF NOT EXISTS "users"(
    "id" bigserial PRIMARY KEY,
    "username" VARCHAR (127) NOT NULL UNIQUE,
    "password" VARCHAR (127) NOT NULL
)
