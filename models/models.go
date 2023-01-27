package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Recipe struct {
	Id          int64     `json:"id"`
	AccountId   int       `json:"accountId"`
	RecipeName  string    `json:"recipeName"`
	RecipeSteps string    `json:"recipeSteps"`
	CreatedOn   time.Time `json:"createdOn"`
	UpdatedOn   time.Time `json:"updatedOn"`
	// TODO: implement this later.
	// IngredientQuantity []IngredientQuantityType `json:"ingredientQuantity"`
}

type Account struct {
	Id        int       `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
	// Recipes   []Recipe  `json:"recipes"`
}

type User struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
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

func (u *User) hashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return err
	}

	u.Password = string(bytes)
	return nil
}
