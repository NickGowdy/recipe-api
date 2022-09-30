package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

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

	dirname, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirname)

	file, err := ioutil.ReadFile(fmt.Sprintf("%s/scripts/db/init.sql", dirname))
	if err != nil {
		panic(err)
	}
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		tx.Rollback()
	}()
	for _, q := range strings.Split(string(file), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		if err != nil {
			panic(err)
		}
	}
	tx.Commit()
}
