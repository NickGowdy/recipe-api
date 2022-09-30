-- name: create-user-table
CREATE TABLE IF NOT EXISTS account (
  id SERIAL PRIMARY KEY,
  first_name varchar(45) NOT NULL,
  last_name varchar(45) NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);

-- name: create-recipes-table
CREATE TABLE IF NOT EXISTS recipe (
  id SERIAL PRIMARY KEY,
  account_id integer REFERENCES account (id),
  name varchar(45) NOT NULL,
  text varchar NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);