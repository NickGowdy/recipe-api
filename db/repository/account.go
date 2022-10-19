package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/recipe-api/models"
)

func GetAccount(id int) (a models.Account, err error) {

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Panic(err)
	}

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
