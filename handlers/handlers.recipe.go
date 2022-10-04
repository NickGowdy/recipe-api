package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/recipe-api/m/recipe"
)

func handleRecipes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		recipeList := recipe.GetRecipes()
		j, err := json.Marshal(recipeList)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(j)

	}
}
