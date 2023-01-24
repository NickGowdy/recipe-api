package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/recipe-api/recipeDb"
)

type Recipe struct {
	Id          int       `json:"id"`
	AccountId   int       `json:"accountId"`
	RecipeName  string    `json:"recipeName"`
	RecipeSteps string    `json:"recipeSteps"`
	CreatedOn   time.Time `json:"createdOn"`
	UpdatedOn   time.Time `json:"updatedOn"`
	// TODO: implement this later.
	// IngredientQuantity []IngredientQuantityType `json:"ingredientQuantity"`
}

func GetRecipes(db *recipeDb.RecipeDb) (*[]Recipe, error) {

	rows, err := db.SqlDb.Query("select * from recipe")
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	var recipes []Recipe

	for rows.Next() {
		var r Recipe
		err := rows.Scan(
			&r.Id,
			&r.AccountId,
			&r.RecipeName,
			&r.RecipeSteps,
			&r.CreatedOn,
			&r.UpdatedOn)
		if err != nil {
			log.Print(err)
		}

		recipes = append(recipes, r)
		log.Print(recipes)
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	return &recipes, nil
}

func GetRecipe(db *recipeDb.RecipeDb, recipeId int) (*Recipe, error) {

	row := db.SqlDb.QueryRow("SELECT * FROM recipe WHERE id=$1", recipeId)
	var recipe Recipe

	switch err := row.Scan(
		&recipe.Id,
		&recipe.AccountId,
		&recipe.RecipeName,
		&recipe.RecipeSteps,
		&recipe.CreatedOn,
		&recipe.UpdatedOn,
	); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return &recipe, nil
	default:
		panic(err)
	}
}

func InsertRecipe(db *recipeDb.RecipeDb, nr *Recipe) (b bool, err error) {
	q := `INSERT INTO recipe ("account_id", "recipe_name", "recipe_steps", "created_on", "updated_on") VALUES($1, $2, $3, now(), now())`
	_, err = db.SqlDb.Exec(q, nr.AccountId, nr.RecipeName, nr.RecipeSteps)

	if err != nil {
		log.Print(err)
	}

	return true, nil
}

func UpdateRecipe(db *recipeDb.RecipeDb, er *Recipe, recipeid int) (d bool, err error) {
	q := `
		UPDATE recipe
		SET recipe_name = $2, recipe_steps = $3
		WHERE id = $1;`

	_, err = db.SqlDb.Exec(q, recipeid, er.RecipeName, er.RecipeSteps)
	if err != nil {
		log.Print(err)
	}

	return true, nil
}

func DeleteRecipe(db *recipeDb.RecipeDb, recipeId int) (d bool, err error) {
	q := `DELETE FROM recipe WHERE id=$1`
	_, err = db.SqlDb.Exec(q, recipeId)

	if err != nil {
		log.Print(err)
	}

	return true, nil
}