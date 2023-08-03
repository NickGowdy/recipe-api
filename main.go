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
	migrate "github.com/rubenv/sql-migrate"

	_ "github.com/lib/pq"
)

const (
	driver        = "postgres"
	migrationsDir = "migrations"
	seconds       = 30
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

	migrations := &migrate.FileMigrationSource{
		Dir: migrationsDir,
	}

	number, err := migrate.Exec(db, driver, migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Applied %d migrations!\n", number)

	queries := database.New(db)

	ctx, cancel := context.WithTimeout(context.Background(), seconds*time.Second)
	defer cancel()

	userRepository := repository.NewUserRepository(queries, &ctx)
	recipeRepository := repository.NewRecipeRepository(queries, &ctx)

	routes := handlers.NewHandlers(&userRepository, &recipeRepository)

	routes.SetupRoutes()
}
