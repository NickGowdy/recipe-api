package main

import (
	"github.com/recipe-api/handlers"
	"github.com/recipe-api/recipeDb"
	"github.com/recipe-api/repository"
)

func main() {
	recipeDb.Migrate()
	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)

	handlers.SetupRoutes(&repo)
}
