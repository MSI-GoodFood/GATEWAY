CREATE SCHEMA IF NOT EXISTS public;
grant usage on schema public to public;
grant create on schema public to public;

CREATE TABLE IF NOT EXISTS "User" (
    "id" uuid PRIMARY KEY,
    "email" varchar UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS "Todo" (
    "id" uuid PRIMARY KEY,
    "text" varchar NOT NULL,
    "done" smallint NOT NULL DEFAULT 0,
    "user_id" uuid NOT NULL
);
