package models

type Account struct {
	Id        int      `json:"id"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Recipes   []Recipe `json:"recipes"`
}

type Recipe struct {
	Id                 int                      `json:"id"`
	UserId             int                      `json:"userId"`
	Name               string                   `json:"name"`
	Text               string                   `json:"Text"`
	IngredientQuantity []IngredientQuantityType `json:"ingredientQuantity"`
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
