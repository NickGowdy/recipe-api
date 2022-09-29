package recipe

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const recipesPath = "recipes"

type Recipe struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func SetupRoutes(apiBasePath string) {
	recipeHandlers := http.HandlerFunc(handleRecipes)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, recipesPath), recipeHandlers)
	http.ListenAndServe(":8080", recipeHandlers)
}

func handleRecipes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		recipeList := getRecipes()
		j, err := json.Marshal(recipeList)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(j)

	}
}

func getRecipes() []Recipe {
	return []Recipe{
		{
			Id:          1,
			Name:        "Chilli Con Carne",
			Description: "Classic Mexican dish that's pure comfort food",
		},
	}
}
