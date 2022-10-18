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
  id INTEGER PRIMARY KEY generated always as identity,
  account_id integer REFERENCES account (id),
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

--name: insert-account
INSERT INTO
  account (
    first_name,
    last_name,
    created_on,
    updated_on
  )
SELECT
  'Nick',
  'Gowdy',
  now(),
  now()
WHERE
  NOT EXISTS (
    SELECT
      id
    FROM
      account
    WHERE
      id = 1
  );

--name: insert-recipe
INSERT INTO
  recipe (
    account_id,
    recipe_name,
    recipe_steps,
    created_on,
    updated_on
  )
SELECT
  1,
  'Lasagna',
  'Some steps to make Lasagna',
  now(),
  now()
WHERE
  NOT EXISTS (
    SELECT
      id
    FROM
      recipe
    WHERE
      id = 1
  );