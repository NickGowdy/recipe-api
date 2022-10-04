package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/recipe-api/models"
)

func GetUser(id int) (returnedAccount models.User, err error) {

	var psqlconn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Panic(err)
	}

	rows, err := db.Query("select id, name from users where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&returnedAccount)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return returnedAccount, nil
}
