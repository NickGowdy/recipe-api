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

func main() {
	setupRoutes()
}

func setupRoutes() {
	connectToDb()

	recipe.SetupRoutes(basePath)
}

func connectToDb() {

	var psqlconn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

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

func migrations(db *sql.DB) {
	dirname, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirname)

	dot, err := dotsql.LoadFromFile(fmt.Sprintf("%s/scripts/db/init.sql", dirname))

	if err != nil {
		log.Fatal(err)
	}

	runScript(db, dot, "create-user-table")
	runScript(db, dot, "create-recipes-table")
}

func runScript(db *sql.DB, dot *dotsql.DotSql, name string) {
	_, err := dot.Exec(db, name)
	if err != nil {
		log.Fatal(err)
	}
}
