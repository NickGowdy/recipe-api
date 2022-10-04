package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/recipe-api/models"
)

func GetAccount(id int) (returnedAccount models.Account, err error) {

	var psqlconn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
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
			&returnedAccount.Id,
			&returnedAccount.Firstname,
			&returnedAccount.Lastname,
			&returnedAccount.CreatedDate,
			&returnedAccount.UpdatedDate)
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
