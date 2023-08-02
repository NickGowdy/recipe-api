package repository

import (
	"context"
	"log"

	"github.com/recipe-api/database"
)

type UserRepository struct {
	queries *database.Queries
	context *context.Context
}

func NewUserRepository(queries *database.Queries, context *context.Context) UserRepository {
	return UserRepository{
		queries: queries,
		context: context,
	}
}

func (ur *UserRepository) InsertRecipeUser(firstname string, lastname string, email string, hashedPwd string) (int64, error) {
	user, err := ur.queries.CreateRecipeUser(*ur.context, database.CreateRecipeUserParams{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  hashedPwd,
	})

	if err != nil {
		log.Print(err)
		return 0, err
	}

	return int64(user.ID), err
}

func (ur *UserRepository) DeleteRecipeUser(recipeUserId int) (bool, error) {

	err := ur.queries.DeleteRecipeUser(*ur.context, int32(recipeUserId))

	if err != nil {
		log.Print(err)
		return false, err
	}
	return true, err
}

func (r *UserRepository) GetRecipeUserPwd(email string) (*database.RecipeUser, error) {

	user, err := r.queries.GetRecipeUserPwd(*r.context, email)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &user, err
}
