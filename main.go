package main

import (
	"github.com/joho/godotenv"
	"github.com/recipe-api/handlers"
	"github.com/recipe-api/recipeDb"
	"github.com/recipe-api/repository"
)

func main() {
	godotenv.Load()

	recipeDb.Migrate()
	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)

	handlers.SetupRoutes(&repo)
}
