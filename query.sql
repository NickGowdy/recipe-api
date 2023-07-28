-- name: GetRecipe :one
SELECT
    *
FROM
    recipe
WHERE
    id = $1
    AND recipe_user_id = $2
LIMIT
    1;

-- name: ListRecipes :many
SELECT
    *
FROM
    recipe
WHERE
    recipe_user_id = $1
ORDER BY
    recipe_name;

-- name: CreateRecipe :one
INSERT INTO
    recipe (
        recipe_user_id,
        recipe_name,
        recipe_steps,
        created_on,
        updated_on
    )
VALUES
    ($1, $2, $3, now(), now()) RETURNING *;

-- name: UpdateRecipe :exec
UPDATE
    recipe 
        SET
            recipe_name = $3,
            recipe_steps = $4
        WHERE
            id = $1
            AND recipe_user_id = $2;

-- name: DeleteRecipe :exec
DELETE FROM
    recipe
WHERE
    id = $1
    AND recipe_user_id = $2;

-- name: CreateRecipeUser :one
INSERT INTO
    recipe_user (
        first_name,
        last_name,
        email,
        password,
        created_on,
        updated_on
    )
VALUES
    ($1, $2, $3, $4, now(), now()) RETURNING *;

-- name: DeleteRecipeUser :exec
DELETE FROM
    recipe_user
WHERE
    id = $1;

-- name: GetRecipeUserPwd :one
SELECT
    *
FROM
    recipe_user
WHERE
    email = $1
LIMIT
    1;