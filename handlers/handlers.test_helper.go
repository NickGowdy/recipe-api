package handlers

import (
	"os"

	"github.com/recipe-api/recipeDb"
	"github.com/recipe-api/repository"
)

func SetupEnv() {
	os.Setenv("user", "postgres")
	os.Setenv("password", "postgres")
	os.Setenv("dbname", "recipes_db")
	os.Setenv("host", "localhost")
	os.Setenv("port", "5432")
}

func Teardown(recipeUserId int) {
	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)
	repo.DeleteRecipeUser(recipeUserId)
}
