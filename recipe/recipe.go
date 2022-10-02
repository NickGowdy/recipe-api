package recipe

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/recipe-api/models"
)

func HandleRecipes(w http.ResponseWriter, r *http.Request) {
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
			Id:   1,
			Name: "Chilli Con Carne",
			Text: "Classic Mexican dish that's pure comfort food",
			IngredientQuantity: []models.IngredientQuantityType{
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
