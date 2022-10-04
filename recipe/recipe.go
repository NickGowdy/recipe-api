package recipe

import "github.com/recipe-api/m/models"

func GetRecipes() []models.Recipe {
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
