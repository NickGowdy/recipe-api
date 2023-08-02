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
	"github.com/recipe-api/recipeDb"
	"github.com/recipe-api/repository"

	_ "github.com/lib/pq" // <------------ here
)

const (
	driver  = "postgres"
	seconds = 30
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

	// need this for now.....
	recipeDb.Migrate()

	queries := database.New(db)

	ctx, cancel := context.WithTimeout(context.Background(), seconds*time.Second)
	defer cancel()

	userRepository := repository.NewUserRepository(queries, &ctx)
	recipeRepository := repository.NewRecipeRepository(queries, &ctx)

	routes := handlers.NewHandlers(&userRepository, &recipeRepository)

	routes.SetupRoutes()
}
