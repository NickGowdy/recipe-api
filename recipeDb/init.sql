-- name: create-recipe_user-table
CREATE TABLE IF NOT EXISTS recipe_user (
  id INTEGER PRIMARY KEY generated always as identity,
  first_name VARCHAR(45) NOT NULL,
  last_name VARCHAR(45) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password TEXT NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);

-- name: create-recipe-table
CREATE TABLE IF NOT EXISTS recipe (
  id INTEGER PRIMARY KEY generated always as identity,
  recipe_user_id integer REFERENCES recipe_user (id),
  recipe_name VARCHAR(45) NOT NULL,
  recipe_steps VARCHAR NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);

-- name: create-ingredient-table
CREATE TABLE IF NOT EXISTS ingredient (
  id INTEGER PRIMARY KEY generated always as identity,
  name VARCHAR(45) NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);

-- name: create-quantity_type-table
CREATE TABLE IF NOT EXISTS quantity_type (
  id INTEGER PRIMARY KEY generated always as identity,
  type VARCHAR(45) NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);

-- name: create-ingredient_quantity_type-table
CREATE TABLE IF NOT EXISTS ingredient_quantity_type (
  id INTEGER PRIMARY KEY generated always as identity,
  ingredient_id INTEGER REFERENCES ingredient (id),
  quantity_type_id INTEGER REFERENCES quantity_type (id),
  amount INTEGER NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);