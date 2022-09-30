package recipe

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/recipe-api/m/models"
)

const recipesPath = "recipes"

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

func getRecipes() []models.Recipe {
	return []models.Recipe{
		{
			Id:          1,
			Name:        "Chilli Con Carne",
			Description: "Classic Mexican dish that's pure comfort food",
			IngredientQuantity: []models.IngredientQuantity{
				{
					Id:             1,
					IngredientId:   1,
					QuantityTypeId: 1,
					QuantityType: models.QuantityType{
						Id:   1,
						Type: "Default",
					},
					Ingredient: models.Ingredient{
						Id:   1,
						Name: "Onion",
					},
					Amount: 1,
				},
			},
		},
	}
}
