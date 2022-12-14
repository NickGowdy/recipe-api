package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
)

func Migrate() {
	db := connect()
	dot := getDirectory()

	fmt.Println(os.Getenv("APP_ENV"))
	fmt.Println("Running migrations")
	runScript(db, dot, "create-account-table")
	runScript(db, dot, "create-recipe-table")
	runScript(db, dot, "create-ingredient-table")
	runScript(db, dot, "create-quantity_type-table")
	runScript(db, dot, "create-ingredient_quantity_type-table")

	// Seed data for dev environment
	if os.Getenv("APP_ENV") == "development" {
		fmt.Println("Seeding dev data")
		runScript(db, dot, "insert-account")
		runScript(db, dot, "insert-recipe")
	}

	// close database
	defer db.Close()
}

func connect() *sql.DB {
	var psqlconn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// check db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
	return db
}

func getDirectory() *dotsql.DotSql {
	dirname, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirname)

	dot, err := dotsql.LoadFromFile(fmt.Sprintf("%s/scripts/db/init.sql", dirname))

	if err != nil {
		log.Fatal(err)
	}

	return dot
}

func runScript(db *sql.DB, dot *dotsql.DotSql, name string) {
	_, err := dot.Exec(db, name)
	if err != nil {
		log.Fatal(err)
	}
}
