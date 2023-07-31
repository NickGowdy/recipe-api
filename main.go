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
)

func main() {
	godotenv.Load()

	recipeDb.Migrate()

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Panic(err)
	}

	queries := database.New(db)

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	repo := repository.NewRecipeRepository(queries, &ctxTimeout)

	handlers.SetupRoutes(&repo)
}
