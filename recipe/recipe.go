package recipe

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/recipe-api/models"
)

func GetRecipes(recipeId int) (returnedRecipes []models.Recipe, err error) {
	var psqlconn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Panic(err)
	}

	rows, err := db.Query("SELECT * FROM recipe where account_id=$1", recipeId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rows)

	for rows.Next() {
		var recipe models.Recipe
		err := rows.Scan(
			&recipe.Id,
			&recipe.AccountId,
			&recipe.RecipeName,
			&recipe.RecipeSteps,
			&recipe.CreatedDate,
			&recipe.UpdatedDate,
		)
		if err != nil {
			log.Fatal(err)
		}
		returnedRecipes = append(returnedRecipes, recipe)
	}
	return returnedRecipes, nil
}
