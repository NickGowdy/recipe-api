-- +migrate Up
CREATE TABLE IF NOT EXISTS recipe_user (
  id INTEGER PRIMARY KEY generated always as identity,
  first_name VARCHAR(45) NOT NULL,
  last_name VARCHAR(45) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password TEXT NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL,
  UNIQUE(email)
);

-- +migrate Down
DROP TABLE recipe_user;