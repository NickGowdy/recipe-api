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
		os.Getenv("host"), os.Getenv("dbport"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

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

	log.Println("Loading routes...")
	mr := mux.NewRouter()

	mr.HandleFunc("/register", handlers.PostRegisterHandler(&userRepository)).Methods("POST")
	mr.HandleFunc("/login", handlers.PostLoginHandler(&userRepository)).Methods("POST")

	mr.Handle("/recipe", handlers.Middleware(handlers.GetAllRecipesHandler(&recipeRepository))).Methods("GET")
	mr.Handle("/recipe/{id}", handlers.Middleware(handlers.GetRecipeHandler(&recipeRepository))).Methods("GET")
	mr.Handle("/recipe", handlers.Middleware(handlers.InsertRecipeHandler(&recipeRepository))).Methods("POST")
	mr.Handle("/recipe/{id}", handlers.Middleware(handlers.UpdateRecipeHandler(&recipeRepository))).Methods("PUT")
	mr.Handle("/recipe/{id}", handlers.Middleware(handlers.DeleteRecipeHandler(&recipeRepository))).Methods("DELETE")

	mr.HandleFunc("/health-check", HealthCheck).Methods("GET")

	http.Handle("api/", mr)
	serverPort := fmt.Sprintf(":%s", os.Getenv("serverport"))
	err = http.ListenAndServe(serverPort, mr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("API is now up...")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}
