package models

type Recipe struct {
	Id                 int                  `json:"id"`
	Name               string               `json:"name"`
	Text               string               `json:"Text"`
	Instruction        Instruction          `json:"instruction"`
	IngredientQuantity []IngredientQuantity `json:"ingredientQuantity"`
}

type Instruction struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

type Ingredient struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type QuantityType struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

type IngredientQuantity struct {
	Id             int          `json:"id"`
	IngredientId   int          `json:"ingredientId"`
	QuantityTypeId int          `json:"quantityTypeId"`
	Amount         int          `json:"quantity"`
	Ingredient     Ingredient   `json:"ingredient"`
	QuantityType   QuantityType `json:"quantityType"`
}
