package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/recipe-api/m/recipe"
)

const basePath = "/api"

const (
	host     = "database"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "recipes"
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
}
