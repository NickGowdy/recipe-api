package handlers

import (
	"fmt"
	"net/http"
)

const recipesPath = "recipes"

func SetupRoutes(apiBasePath string) {
	recipeHandlers := http.HandlerFunc(handleRecipes)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, recipesPath), recipeHandlers)
}
