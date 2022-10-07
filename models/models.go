package models

import "time"

type Account struct {
	Id          int       `json:"id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	Recipes     []Recipe  `json:"recipes"`
}

type Recipe struct {
	Id          int       `json:"id"`
	AccountId   int       `json:"accountId"`
	RecipeName  string    `json:"recipeName"`
	RecipeSteps string    `json:"recipeSteps"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	// TODO: implement this later.
	// IngredientQuantity []IngredientQuantityType `json:"ingredientQuantity"`
}

type Ingredient struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type QuantityType struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

type IngredientQuantityType struct {
	Id             int          `json:"id"`
	IngredientId   int          `json:"ingredientId"`
	QuantityTypeId int          `json:"quantityTypeId"`
	Amount         int          `json:"quantity"`
	Ingredient     Ingredient   `json:"ingredient"`
	QuantityType   QuantityType `json:"quantityType"`
}
