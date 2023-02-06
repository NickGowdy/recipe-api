package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/recipe-api/models"
	"github.com/recipe-api/recipeDb"
	"github.com/recipe-api/repository"
)

type TestUser struct {
	email    string
	password string
}

func SetupEnvVars() {
	os.Setenv("user", "postgres")
	os.Setenv("password", "postgres")
	os.Setenv("dbname", "recipes_db")
	os.Setenv("host", "localhost")
	os.Setenv("port", "5432")
}

func SetupUser() *TestUser {
	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)
	email := fmt.Sprintf("%s@gmail.com", randomString(12))
	password := randomString(12)

	body, _ := json.Marshal(models.Register{
		Firstname: randomString(12),
		Lastname:  randomString(12),
		Email:     email,
		Password:  randomString(12),
	})
	req, _ := http.NewRequest("POST", "/register", bytes.NewReader(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostRegisterHandler(&repo))

	handler.ServeHTTP(rr, req)

	var recipeUserId int
	json.NewDecoder(rr.Body).Decode(&recipeUserId)
	return &TestUser{email: email, password: password}
}

func SetupRecipe(bearer string) int64 {
	recipeToInsert := models.Recipe{
		RecipeName:  "Nick's recipe",
		RecipeSteps: "Some steps for Nick's recipe",
	}
	body, _ := json.Marshal(recipeToInsert)
	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)

	req, err := http.NewRequest("POST", "/recipe", bytes.NewReader(body))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	if err != nil {
		log.Panic(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(Middleware(InsertRecipeHandler(&repo)))

	handler.ServeHTTP(rr, req)

	var recipeId int64

	json.NewDecoder(rr.Body).Decode(&recipeId)
	return recipeId
}

func SetupToken(tu *TestUser) string {
	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)

	body, _ := json.Marshal(models.Credentials{
		Email:    tu.email,
		Password: tu.password,
	})

	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostLoginHandler(&repo))
	handler.ServeHTTP(rr, req)

	bodyBytes, _ := io.ReadAll(rr.Body)
	return string(bodyBytes)
}

func Teardown(recipeUserId int) {
	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)
	repo.DeleteRecipeUser(recipeUserId)
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}
