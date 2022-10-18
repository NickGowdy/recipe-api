package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/recipe-api/models"
)

func GetRecipes(accountId int) (rs []models.Recipe, err error) {
	db := getConnection()

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
	db := getConnection()

	rows, err := db.Query("select * from recipe where id=$1 and account_id=$2", id, account_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&r.Id,
			&r.AccountId,
			&r.RecipeName,
			&r.RecipeSteps,
			&r.CreatedOn,
			&r.UpdatedOn)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return r, nil
}

func SaveRecipe(nr *models.Recipe) (b bool, err error) {
	db := getConnection()

	q := `INSERT INTO recipe ("account_id", "recipe_name", "recipe_steps", "created_on", "updated_on") VALUES($1, $2, $3, now(), now())`
	_, err = db.Exec(q, nr.AccountId, nr.RecipeName, nr.RecipeSteps)

	if err != nil {
		log.Panic(err)
	}

	return true, nil
}

func DeleteRecipe(recipeId int, accountid int) (d bool, err error) {
	db := getConnection()

	q := `DELETE FROM recipe WHERE id=$1 AND account_id=$2`
	_, err = db.Exec(q, recipeId, accountid)

	if err != nil {
		log.Panic(err)
	}

	return true, nil
}

func UpdateRecipe(er *models.Recipe, recipeid int, accountid int) (d bool) {
	db := getConnection()

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

func getConnection() sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Panic(err)
	}

	return *db
}
