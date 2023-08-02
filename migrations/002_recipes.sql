-- +migrate Up
CREATE TABLE IF NOT EXISTS recipe (
  id INTEGER PRIMARY KEY generated always as identity,
  recipe_user_id integer REFERENCES recipe_user (id) NOT NULL,
  recipe_name VARCHAR(45) NOT NULL,
  recipe_steps VARCHAR NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS ingredient (
  id INTEGER PRIMARY KEY generated always as identity,
  name VARCHAR(45) NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS quantity_type (
  id INTEGER PRIMARY KEY generated always as identity,
  type VARCHAR(45) NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS ingredient_quantity_type (
  id INTEGER PRIMARY KEY generated always as identity,
  ingredient_id INTEGER REFERENCES ingredient (id),
  quantity_type_id INTEGER REFERENCES quantity_type (id),
  amount INTEGER NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
);

-- +migrate Down
DROP TABLE ingredient_quantity_type;
DROP TABLE quantity_type;
DROP TABLE ingredient;