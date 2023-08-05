package testSetup

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/recipe-api/database"
	"github.com/recipe-api/repository"
	"github.com/recipe-api/security"
)

const (
	firstname     string = "Test"
	lastname      string = "User"
	email         string = "testuser@gmail.com"
	password      string = "password"
	driver               = "postgres"
	migrationsDir        = "migrations"
	seconds              = 30
)

// func SetupEnvVars() {
// 	os.Setenv("user", "postgres")
// 	os.Setenv("password", "postgres")
// 	os.Setenv("dbname", "recipes_db")
// 	os.Setenv("host", "localhost")
// 	os.Setenv("dbport", "5432")
// }

func GetTestToken() string {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("dbport"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, _ := sql.Open(driver, psqlconn)
	defer db.Close()

	queries := database.New(db)

	ctx, cancel := context.WithTimeout(context.Background(), seconds*time.Second)
	defer cancel()

	repo := repository.NewUserRepository(queries, &ctx)

	userId, _ := repo.InsertRecipeUser(firstname, lastname, email, password)

	token, _ := security.GenerateToken(userId)
	return token
}

func Teardown(recipeUserId int) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("dbport"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, _ := sql.Open(driver, psqlconn)
	defer db.Close()

	queries := database.New(db)

	ctx, cancel := context.WithTimeout(context.Background(), seconds*time.Second)
	defer cancel()

	repo := repository.NewUserRepository(queries, &ctx)
	repo.DeleteRecipeUser(recipeUserId)
}
