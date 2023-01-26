package main

import (
	"github.com/recipe-api/handlers"
	"github.com/recipe-api/recipeDb"
)

var db *recipeDb.RecipeDb

func main() {
	recipeDb.Migrate()
	db = recipeDb.NewRecipeDb()

	handlers.SetupRoutes(db)
}
