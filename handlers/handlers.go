package handlers

import (
	"fmt"
	"net/http"

	"github.com/recipe-api/m/recipe"
)

const recipesPath = "recipes"

// const userPath = "accounts"

func SetupRoutes(apiBasePath string) {
	// accountHandlers := http.HandlerFunc(account.HandleAccounts)
	recipeHandlers := http.HandlerFunc(recipe.HandleRecipes)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, recipesPath), recipeHandlers)
	http.ListenAndServe(":8080", recipeHandlers)

	// fmt.Println(accountHandlers)
}
