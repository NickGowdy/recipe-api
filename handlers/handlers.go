package handlers

import (
	"fmt"
	"net/http"
)

const recipesPath = "recipes"
const accountsPath = "accounts"

func SetupRoutes(apiBasePath string) {
	recipeHandlers := http.HandlerFunc(handleRecipes)
	accountHandler := http.HandlerFunc(handleAccount)

	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, recipesPath), recipeHandlers)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, accountsPath), accountHandler)
}
