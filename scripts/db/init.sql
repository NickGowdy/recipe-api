-- name: create-recipes-table
CREATE TABLE IF NOT EXISTS recipe (
  id SERIAL PRIMARY KEY,
  name varchar(45) NOT NULL,
  text varchar NOT NULL,
  created_on TIMESTAMP NOT NULL,
  updated_on TIMESTAMP NOT NULL
)