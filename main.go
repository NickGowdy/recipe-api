package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/recipe-api/database"
	"github.com/recipe-api/handlers"
	"github.com/recipe-api/repository"
)

const (
	driver  = "postgres"
	timeout = 30
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, err := sql.Open(driver, psqlconn)

	if err != nil {
		log.Fatal(err)
	}

	queries := database.New(db)

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	repo := repository.NewRecipeRepository(queries, &ctx)

	handlers.SetupRoutes(&repo)
}
