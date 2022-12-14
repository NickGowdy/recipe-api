package repository

import (
	"database/sql"
	"log"

	"github.com/recipe-api/models"
)

func GetRecipes(accountId int) (rs []models.Recipe, err error) {
	db := Database()

	rows, err := db.Query("select * from recipe where account_id=$1", accountId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Recipe
		err := rows.Scan(
			&r.Id,
			&r.AccountId,
			&r.RecipeName,
			&r.RecipeSteps)
		if err != nil {
			log.Fatal(err)
		}

		rs = append(rs, r)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return rs, nil
}

func GetRecipe(id int, account_id int) (r models.Recipe, err error) {
	db := Database()

	row := db.QueryRow("SELECT * FROM recipe WHERE id=$1 and account_id=$2", id, account_id)

	switch err := row.Scan(&r.Id, &r.AccountId, &r.RecipeName, &r.RecipeSteps, &r.CreatedOn, &r.UpdatedOn); err {
	case sql.ErrNoRows:
		return r, err
	case nil:
		return r, nil
	default:
		panic(err)
	}
}

func SaveRecipe(nr *models.Recipe) (b bool, err error) {
	db := Database()

	q := `INSERT INTO recipe ("account_id", "recipe_name", "recipe_steps", "created_on", "updated_on") VALUES($1, $2, $3, now(), now())`
	_, err = db.Exec(q, nr.AccountId, nr.RecipeName, nr.RecipeSteps)

	if err != nil {
		log.Panic(err)
	}

	return true, nil
}

func DeleteRecipe(recipeId int, accountid int) (d bool, err error) {
	db := Database()

	q := `DELETE FROM recipe WHERE id=$1 AND account_id=$2`
	_, err = db.Exec(q, recipeId, accountid)

	if err != nil {
		log.Panic(err)
	}

	return true, nil
}

func UpdateRecipe(er *models.Recipe, recipeid int, accountid int) (d bool) {
	db := Database()

	dr, _ := GetRecipe(recipeid, accountid)

	if dr.Id != recipeid && dr.AccountId != accountid {
		return false
	}

	q := `
		UPDATE recipe
		SET recipe_name = $3, recipe_steps = $4
		WHERE id = $1 AND account_id = $2;`

	_, err := db.Exec(q, recipeid, accountid, er.RecipeName, er.RecipeSteps)
	if err != nil {
		panic(err)
	}

	return true
}
