package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/recipe-api/recipe"
)

func handleRecipes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		recipeList, err := recipe.GetRecipes(1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		j, err := json.Marshal(recipeList)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(j)
	}
}
