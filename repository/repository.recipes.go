package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/recipe-api/database"
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
	var parent sql.NullInt32

	if err := parent.Scan(&recipeUserId); err != nil {
		log.Print(err)
		return nil, err
	}

	recipes, err := r.queries.ListRecipes(*r.context, parent)
	if err != nil {
		log.Print(err)
	}

	return recipes, err

}

// func (r *RecipeRepository) GetRecipe(recipeId int, recipeUserid int) (*models.Recipe, error) {

// 	row := r.db.SqlDb.QueryRow("SELECT * FROM recipe WHERE id=$1 AND recipe_user_id=$2", recipeId, recipeUserid)
// 	var recipe models.Recipe

// 	switch err := row.Scan(
// 		&recipe.Id,
// 		&recipe.RecipeUserId,
// 		&recipe.RecipeName,
// 		&recipe.RecipeSteps,
// 		&recipe.CreatedOn,
// 		&recipe.UpdatedOn,
// 	); err {
// 	case sql.ErrNoRows:
// 		return nil, err
// 	case nil:
// 		return &recipe, nil
// 	default:
// 		panic(err)
// 	}
// }

// func (r *RecipeRepository) InsertRecipe(recipeUserId int, ir *models.SaveRecipe) (b int64, err error) {
// 	var id int64
// 	var cols = "(recipe_user_id, recipe_name, recipe_steps, created_on, updated_on)"
// 	var values = "($1, $2, $3, now(), now())"

// 	var query = fmt.Sprintf(
// 		"INSERT INTO recipe %s VALUES %s RETURNING id",
// 		cols, values,
// 	)

// 	if err := r.db.SqlDb.QueryRow(
// 		query,
// 		recipeUserId, ir.RecipeName, ir.RecipeSteps,
// 	).Scan(&id); err != nil {
// 		panic(err)
// 	}

// 	if err != nil {
// 		log.Print(err)
// 		return 0, err
// 	}

// 	return id, nil
// }

// func (r *RecipeRepository) UpdateRecipe(recipeid int, recipeUserId int, recipe *models.SaveRecipe) (d bool, err error) {
// 	q := `
// 		UPDATE recipe
// 		SET recipe_name = $3, recipe_steps = $4
// 		WHERE id = $1 AND recipe_user_id = $2;`

// 	_, err = r.db.SqlDb.Exec(q, recipeid, recipeUserId, recipe.RecipeName, recipe.RecipeSteps)
// 	if err != nil {
// 		log.Print(err)
// 	}

// 	return true, nil
// }

// func (r *RecipeRepository) DeleteRecipe(recipeId int, recipeUserId int) (d bool, err error) {
// 	q := `DELETE FROM recipe WHERE id=$1 AND recipe_user_id=$2;`
// 	_, err = r.db.SqlDb.Exec(q, recipeId, recipeUserId)

// 	if err != nil {
// 		log.Print(err)
// 	}

// 	return true, nil
// }

// func (r *RecipeRepository) InsertRecipeUser(firstname string, lastname string, email string, hashedPwd string) (b int64, err error) {
// 	var id int64
// 	var cols = "(first_name, last_name, email, password, created_on, updated_on)"
// 	var values = "($1, $2, $3, $4, now(), now())"

// 	var query = fmt.Sprintf(
// 		"INSERT INTO recipe_user %s VALUES %s RETURNING id",
// 		cols, values,
// 	)

// 	if err := r.db.SqlDb.QueryRow(
// 		query,
// 		firstname, lastname, email, hashedPwd,
// 	).Scan(&id); err != nil {
// 		log.Print(err)
// 		return 0, err
// 	}

// 	return id, nil
// }

// func (r *RecipeRepository) DeleteRecipeUser(recipeUserId int) (d bool, err error) {
// 	q := "DELETE FROM recipe_user WHERE id=$1;"
// 	_, err = r.db.SqlDb.Exec(q, recipeUserId)

// 	if err != nil {
// 		log.Print(err)
// 	}

// 	return true, nil
// }

// func (r *RecipeRepository) GetRecipeUserPwd(email string) (*models.RecipeUser, error) {

// 	row := r.db.SqlDb.QueryRow("SELECT * FROM recipe_user WHERE email=$1", email)

// 	var ru models.RecipeUser

// 	switch err := row.Scan(
// 		&ru.Id,
// 		&ru.Firstname,
// 		&ru.Password,
// 		&ru.Email,
// 		&ru.Password,
// 		&ru.CreatedOn,
// 		&ru.UpdatedOn,
// 	); err {
// 	case sql.ErrNoRows:
// 		return nil, nil
// 	case nil:
// 		return &ru, nil
// 	default:
// 		panic(err)
// 	}
// }
