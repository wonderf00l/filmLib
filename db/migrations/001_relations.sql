CREATE SCHEMA IF NOT EXISTS filmLib;

SET search_path TO filmLib;

CREATE TABLE IF NOT EXISTS role (
	id smallserial PRIMARY KEY,
	name varchar(100) NOT NULL CHECK (LENGTH(name) >= 1),
	CONSTRAINT role_name_uniq UNIQUE (name)
);
INSERT INTO role (name) VALUES ('admin');
INSERT INTO role (name) VALUES ('regular');


CREATE TABLE IF NOT EXISTS profile (
	id serial PRIMARY KEY,
	username text NOT NULL,
	password text NOT NULL,
    profile_role smallserial NOT NULL,
	created_at timestamptz NOT NULL DEFAULT now(),
	updated_at timestamptz NOT NULL DEFAULT now(),
	deleted_at timestamptz,
	CONSTRAINT profile_username_uniq UNIQUE (username),
    FOREIGN KEY (profile_role) REFERENCES role (id) ON DELETE CASCADE
);

CREATE TYPE sex AS ENUM ('male', 'female');

CREATE TABLE IF NOT EXISTS actor (
	id serial PRIMARY KEY,
	name text NOT NULL,
	gender sex NOT NULL,
	date_of_birth date,
	created_at timestamptz NOT NULL DEFAULT now(),
	updated_at timestamptz NOT NULL DEFAULT now(),
	deleted_at timestamptz
);

CREATE TABLE IF NOT EXISTS film (
	id serial PRIMARY KEY,
	name varchar(150) NOT NULL CHECK (LENGTH(name) >= 1),
	description varchar(1000),
	release_date date,
	rating smallint CHECK (rating BETWEEN 0 and 10),
    actors TEXT[],
	created_at timestamptz NOT NULL DEFAULT now(),
	updated_at timestamptz NOT NULL DEFAULT now(),
	deleted_at timestamptz
);

CREATE EXTENSION IF NOT EXISTS moddatetime;

CREATE OR REPLACE TRIGGER modify_profile_updated_at
	BEFORE UPDATE
	ON profile
	FOR EACH ROW
EXECUTE PROCEDURE moddatetime(updated_at);

CREATE OR REPLACE TRIGGER modify_actor_updated_at
	BEFORE UPDATE
	ON actor
	FOR EACH ROW
EXECUTE PROCEDURE moddatetime(updated_at);

CREATE OR REPLACE TRIGGER modify_film_updated_at
	BEFORE UPDATE
	ON film
	FOR EACH ROW
EXECUTE PROCEDURE moddatetime(updated_at);

