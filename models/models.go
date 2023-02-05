package models

import (
	"time"
)

type Recipe struct {
	Id           int64     `json:"id"`
	RecipeUserId int64     `json:"recipeUserId"`
	RecipeName   string    `json:"recipeName"`
	RecipeSteps  string    `json:"recipeSteps"`
	CreatedOn    time.Time `json:"createdOn"`
	UpdatedOn    time.Time `json:"updatedOn"`
	// TODO: implement this later.
	// IngredientQuantity []IngredientQuantityType `json:"ingredientQuantity"`
}

type RecipeUser struct {
	Id        int64     `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}

type InsertRecipe struct {
	RecipeName  string `json:"recipeName"`
	RecipeSteps string `json:"recipeSteps"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Ingredient struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}

type QuantityType struct {
	Id        int       `json:"id"`
	Type      string    `json:"type"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}

type IngredientQuantityType struct {
	Id             int          `json:"id"`
	IngredientId   int          `json:"ingredientId"`
	QuantityTypeId int          `json:"quantityTypeId"`
	Amount         int          `json:"quantity"`
	Ingredient     Ingredient   `json:"ingredient"`
	QuantityType   QuantityType `json:"quantityType"`
	CreatedOn      time.Time    `json:"createdOn"`
	UpdatedOn      time.Time    `json:"updatedOn"`
}
