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

	rows, err := db.Query("select * from account where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&a.Id,
			&a.Firstname,
			&a.Lastname,
			&a.CreatedDate,
			&a.UpdatedDate)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return a, nil
}
