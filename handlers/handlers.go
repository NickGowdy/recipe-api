package handlers

import (
	"fmt"
	"net/http"

	"github.com/recipe-api/recipe"
	"github.com/recipe-api/user"
)

const recipesPath = "recipes"

const accountsPath = "accounts"

func SetupRoutes(apiBasePath string) {
	userHandler := http.HandlerFunc(user.HandleUser)
	recipeHandlers := http.HandlerFunc(recipe.HandleRecipes)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, recipesPath), recipeHandlers)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, accountsPath), userHandler)
	http.ListenAndServe(":8080", recipeHandlers)
}
