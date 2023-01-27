package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/recipe-api/models"
	"github.com/recipe-api/recipeDb"
)

type Repository struct {
	db *recipeDb.RecipeDb
}

func NewRepository(db *recipeDb.RecipeDb) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) GetRecipes() (*[]models.Recipe, error) {

	rows, err := r.db.SqlDb.Query("select * from recipe")
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	var recipes []models.Recipe

	for rows.Next() {
		var r models.Recipe
		err := rows.Scan(
			&r.Id,
			&r.AccountId,
			&r.RecipeName,
			&r.RecipeSteps,
			&r.CreatedOn,
			&r.UpdatedOn)
		if err != nil {
			log.Print(err)
		}

		recipes = append(recipes, r)
		log.Print(recipes)
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	return &recipes, nil
}

func (r Repository) GetRecipe(recipeId int) (*models.Recipe, error) {

	row := r.db.SqlDb.QueryRow("SELECT * FROM recipe WHERE id=$1", recipeId)
	var recipe models.Recipe

	switch err := row.Scan(
		&recipe.Id,
		&recipe.AccountId,
		&recipe.RecipeName,
		&recipe.RecipeSteps,
		&recipe.CreatedOn,
		&recipe.UpdatedOn,
	); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return &recipe, nil
	default:
		panic(err)
	}
}

func (r Repository) InsertRecipe(nr *models.Recipe) (b int64, err error) {
	var id int64
	var cols = "(account_id, recipe_name, recipe_steps, created_on, updated_on)"
	var values = "($1, $2, $3, now(), now())"

	var query = fmt.Sprintf(
		"INSERT INTO recipe %s VALUES %s RETURNING id",
		cols, values,
	)

	if err := r.db.SqlDb.QueryRow(
		query,
		nr.AccountId, nr.RecipeName, nr.RecipeSteps,
	).Scan(&id); err != nil {
		panic(err)
	}

	if err != nil {
		log.Print(err)
		return 0, err
	}

	return id, nil
}

func (r Repository) UpdateRecipe(recipe *models.Recipe, recipeid int) (d bool, err error) {
	q := `
		UPDATE recipe
		SET recipe_name = $2, recipe_steps = $3
		WHERE id = $1;`

	_, err = r.db.SqlDb.Exec(q, recipeid, recipe.RecipeName, recipe.RecipeSteps)
	if err != nil {
		log.Print(err)
	}

	return true, nil
}

func (r Repository) DeleteRecipe(recipeId int) (d bool, err error) {
	q := `DELETE FROM recipe WHERE id=$1`
	_, err = r.db.SqlDb.Exec(q, recipeId)

	if err != nil {
		log.Print(err)
	}

	return true, nil
}
