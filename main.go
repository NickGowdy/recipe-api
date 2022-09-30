package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
	"github.com/recipe-api/m/recipe"
)

const basePath = "/api"

const (
	host     = "database"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "recipes_db"
)

func main() {
	setupRoutes()
}

func setupRoutes() {
	connectToDb()

	recipe.SetupRoutes(basePath)
}

func connectToDb() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	migrations(db)
}

func migrations(db *sql.DB) (sql.Result, error) {
	dirname, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirname)

	dot, err := dotsql.LoadFromFile(fmt.Sprintf("%s/scripts/db/init.sql", dirname))

	if err != nil {
		log.Fatal(err)
	}

	res, err := dot.Exec(db, "create-recipes-table")
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
