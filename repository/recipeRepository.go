package repository

import (
	"context"
	"log"

	"github.com/recipe-api/database"
	"github.com/recipe-api/models"
)

type RecipeRepository struct {
	queries *database.Queries
	context *context.Context
}

func NewRecipeRepository(queries *database.Queries, context *context.Context) RecipeRepository {
	return RecipeRepository{
		queries: queries,
		context: context,
	}
}

func (r *RecipeRepository) GetRecipes(recipeUserId int) ([]database.Recipe, error) {

	var validRecipeUserId = int32(recipeUserId)

	recipes, err := r.queries.ListRecipes(*r.context, validRecipeUserId)
	if err != nil {
		log.Print(err)
	}

	return recipes, err
}

func (r *RecipeRepository) GetRecipe(recipeId int, recipeUserId int) (*database.Recipe, error) {

	var validRecipeUserId = int32(recipeUserId)

	validRecipeUserId = int32(recipeUserId)

	recipe, err := r.queries.GetRecipe(*r.context, database.GetRecipeParams{ID: int32(recipeId), RecipeUserID: validRecipeUserId})
	return &recipe, err
}

func (r *RecipeRepository) InsertRecipe(recipeUserId int, ir *models.SaveRecipe) (b int32, err error) {

	var validRecipeUserId = int32(recipeUserId)

	newRecipe, err := r.queries.CreateRecipe(*r.context, database.CreateRecipeParams{
		RecipeUserID: validRecipeUserId,
		RecipeName:   ir.RecipeName,
		RecipeSteps:  ir.RecipeSteps,
	})

	return newRecipe.ID, err
}

func (r *RecipeRepository) UpdateRecipe(recipeid int, recipeUserId int, recipe *models.SaveRecipe) (bool, error) {

	var validRecipeUserId = int32(recipeUserId)

	validRecipeUserId = int32(recipeUserId)

	err := r.queries.UpdateRecipe(*r.context, database.UpdateRecipeParams{
		ID:           int32(recipeid),
		RecipeUserID: validRecipeUserId,
		RecipeName:   recipe.RecipeName,
		RecipeSteps:  recipe.RecipeSteps,
	})

	if err != nil {
		log.Print(err)
	}

	return true, err
}

func (r *RecipeRepository) DeleteRecipe(recipeId int, recipeUserId int) (d bool, err error) {
	var validRecipeUserId = int32(recipeUserId)

	validRecipeUserId = int32(recipeUserId)

	r.queries.DeleteRecipe(*r.context, database.DeleteRecipeParams{ID: int32(recipeId), RecipeUserID: validRecipeUserId})

	if err != nil {
		log.Print(err)
	}

	return true, err
}
