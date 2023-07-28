package main

import (
	"github.com/recipe-api/recipeDb"
)

func main() {
	recipeDb.Migrate()
	// db := recipeDb.NewRecipeDb()
	// repo := repository.NewRecipeRepository(db)

	// handlers.SetupRoutes(&repo)
}
