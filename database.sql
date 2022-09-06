CREATE SCHEMA IF NOT EXISTS public;
grant usage on schema public to public;
grant create on schema public to public;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

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

CREATE TABLE IF NOT EXISTS "UserRole"
(
    "id" uuid PRIMARY KEY,
    "label" varchar UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS "Country"
(
    "id" uuid PRIMARY KEY,
    "label" varchar UNIQUE NOT NULL
    "code" varchar UNIQUE NOT NULL
);

INSERT INTO UserRole
VALUES
       (uuid_generate_v4(), "CUSTOMER"),
       (uuid_generate_v4(), "DELIVERY"),
       (uuid_generate_v4(), "ACCOUNTANT"),
       (uuid_generate_v4(), "ADMIN");

INSERT INTO Country VALUES (uuid_generate_v4(), "France", "FR");
