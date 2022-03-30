CREATE SCHEMA IF NOT EXISTS public;
grant usage on schema public to public;
grant create on schema public to public;

CREATE TABLE IF NOT EXISTS "User"
(
    "id" uuid PRIMARY KEY,
    "email" varchar UNIQUE NOT NULL,
    "password" varchar NOT NULL,
    "created_at" timestamp NOT NULL,
    "active" boolean NOT NULL DEFAULT TRUE,
    "id_role" uuid NOT NULL,
    "id_country" uuid NOT NULL
);
