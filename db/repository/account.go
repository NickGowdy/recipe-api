package repository

import (
	"database/sql"

	"github.com/recipe-api/models"
)

func GetAccount(id int) (a models.Account, err error) {

	db := Database()

	row := db.QueryRow("SELECT * FROM account WHERE id=$1", id)

	switch err := row.Scan(&a.Id, &a.Firstname, &a.Lastname, &a.CreatedOn, &a.UpdatedOn); err {
	case sql.ErrNoRows:
		return a, err
	case nil:
		return a, nil
	default:
		panic(err)
	}
}
