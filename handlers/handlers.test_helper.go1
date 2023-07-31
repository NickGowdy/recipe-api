package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/recipe-api/models"
	"github.com/recipe-api/recipeDb"
	"github.com/recipe-api/repository"
)

const (
	firstname string = "Test"
	lastname  string = "User"
	email     string = "testuser@gmail.com"
	password  string = "password"
)

func SetupEnvVars() {
	os.Setenv("user", "postgres")
	os.Setenv("password", "postgres")
	os.Setenv("dbname", "recipes_db")
	os.Setenv("host", "localhost")
	os.Setenv("port", "5432")
}

func GetTestToken() string {
	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)

	body, _ := json.Marshal(models.Register{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Password:  password,
	})
	req, _ := http.NewRequest("POST", "/register", bytes.NewReader(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostRegisterHandler(&repo))

	handler.ServeHTTP(rr, req)

	body, _ = json.Marshal(models.Credentials{
		Email:    email,
		Password: password,
	})

	req, _ = http.NewRequest("POST", "/login", bytes.NewReader(body))

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(PostLoginHandler(&repo))
	handler.ServeHTTP(rr, req)

	bodyBytes, _ := io.ReadAll(rr.Body)
	return string(bodyBytes)
}

func Teardown(recipeUserId int) {
	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)
	repo.DeleteRecipeUser(recipeUserId)
}
