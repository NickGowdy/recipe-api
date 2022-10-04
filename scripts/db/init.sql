-- name: create-account-table
CREATE TABLE IF NOT EXISTS account (
  id INTEGER PRIMARY KEY generated always as identity,
  first_name VARCHAR(45) NOT NULL,
  last_name VARCHAR(45) NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);

-- name: create-recipe-table
CREATE TABLE IF NOT EXISTS recipe (
  id SERIAL PRIMARY KEY,
  account_id integer REFERENCES account (id),
  name VARCHAR(45) NOT NULL,
  text VARCHAR NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);

-- name: create-ingredient-table
CREATE TABLE IF NOT EXISTS ingredient (
  id SERIAL PRIMARY KEY,
  name VARCHAR(45) NOT NULL
);

-- name: create-quantity_type-table
CREATE TABLE IF NOT EXISTS quantity_type (
  id SERIAL PRIMARY KEY,
  type VARCHAR(45) NOT NULL
);

-- name: create-ingredient_quantity_type-table
CREATE TABLE IF NOT EXISTS ingredient_quantity_type (
  id SERIAL PRIMARY KEY,
  ingredient_id INTEGER REFERENCES ingredient (id),
  quantity_type_id INTEGER REFERENCES quantity_type (id),
  amount INTEGER NOT NULL
);

--name: insert-account
INSERT INTO account (first_name, last_name, created_on, updated_on) VALUES('Nick', 'Gowdy', now(), now());