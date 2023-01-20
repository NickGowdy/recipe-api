package repository

import (
	"database/sql"
	"log"
	"time"
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

func GetRecipes() (*[]Recipe, error) {
	db := Database()

	rows, err := db.Query("select * from recipe")
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
			log.Fatal(err)
		}

		recipes = append(recipes, r)
		log.Print(recipes)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &recipes, nil
}

func GetRecipe(recipeId int) (*Recipe, error) {
	db := Database()

	row := db.QueryRow("SELECT * FROM recipe WHERE id=$1", recipeId)
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

// func SaveRecipe(nr *models.Recipe) (b bool, err error) {
// 	db := Database()

// 	q := `INSERT INTO recipe ("account_id", "recipe_name", "recipe_steps", "created_on", "updated_on") VALUES($1, $2, $3, now(), now())`
// 	_, err = db.Exec(q, nr.AccountId, nr.RecipeName, nr.RecipeSteps)

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	return true, nil
// }

// func DeleteRecipe(recipeId int, accountid int) (d bool, err error) {
// 	db := Database()

// 	q := `DELETE FROM recipe WHERE id=$1 AND account_id=$2`
// 	_, err = db.Exec(q, recipeId, accountid)

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	return true, nil
// }

// func UpdateRecipe(er *models.Recipe, recipeid int, accountid int) (d bool) {
// 	db := Database()

// 	dr, _ := GetRecipe(recipeid, accountid)

// 	if dr.Id != recipeid && dr.AccountId != accountid {
// 		return false
// 	}

// 	q := `
// 		UPDATE recipe
// 		SET recipe_name = $3, recipe_steps = $4
// 		WHERE id = $1 AND account_id = $2;`

// 	_, err := db.Exec(q, recipeid, accountid, er.RecipeName, er.RecipeSteps)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return true
// }
