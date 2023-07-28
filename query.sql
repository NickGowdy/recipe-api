-- name: GetRecipe :one
SELECT * FROM recipe
WHERE id = $1 AND recipe_user_id=$2 LIMIT 1;

-- name: ListRecipes :many
SELECT * FROM recipe
WHERE recipe_user_id=$1
ORDER BY recipe_name;
