package main

import (
	"github.com/recipe-api/handlers"
	"github.com/recipe-api/recipeDb"
	"github.com/recipe-api/recipeDb/repository"
)

var db *recipeDb.RecipeDb

func main() {
	recipeDb.Migrate()
	db = recipeDb.NewRecipeDb()
	repo := repository.NewRepository(db)

	handlers.SetupRoutes(&repo)
}
