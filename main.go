package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/recipe-api/database"
	"github.com/recipe-api/handler"
	"github.com/recipe-api/recipe"
	"github.com/recipe-api/repository"
	"github.com/recipe-api/user"
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
		log.Panic(err)
	}

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("dbport"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, err := sql.Open(driver, psqlconn)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	go func() {
		select {
		case <-time.After(seconds * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err()) // prints "context deadline exceeded"
		}
	}()

	userRepository := repository.NewUserRepository(queries, &ctx)
	recipeRepository := repository.NewRecipeRepository(queries, &ctx)

	user := user.NewUser(&userRepository)
	recipe := recipe.NewRecipe(&recipeRepository)

	log.Println("Loading routes...")
	mr := mux.NewRouter()
	h := handler.NewHandler(*user, *recipe)

	h.Start(mr)

	http.Handle("api/", mr)
	serverPort := fmt.Sprintf(":%s", os.Getenv("serverport"))
	err = http.ListenAndServe(serverPort, mr)
	if err != nil {
		log.Fatal(err)
	}
}
