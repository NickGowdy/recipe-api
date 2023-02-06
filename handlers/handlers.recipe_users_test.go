package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/recipe-api/models"
	"github.com/recipe-api/recipeDb"
	"github.com/recipe-api/repository"
)

func TestRegister(t *testing.T) {
	SetupEnvVars()

	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)
	email := fmt.Sprintf("%s@gmail.com", randomString(12))
	password := randomString(12)

	body, _ := json.Marshal(models.Register{
		Firstname: randomString(12),
		Lastname:  randomString(12),
		Email:     email,
		Password:  password,
	})
	req, err := http.NewRequest("POST", "/register/", bytes.NewReader(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostRegisterHandler(&repo))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var recipeUserId int

	err = json.NewDecoder(rr.Body).Decode(&recipeUserId)

	if err != nil {
		t.Error(err)
	}

	if recipeUserId == 0 {
		t.Errorf("recipe user id should be greather than 0 but is: %v", recipeUserId)
	}

	Teardown(recipeUserId)
}

func TestLogin(t *testing.T) {
	SetupEnvVars()

	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)
	email := fmt.Sprintf("%s@gmail.com", randomString(12))
	password := randomString(12)

	body, _ := json.Marshal(models.Register{
		Firstname: randomString(12),
		Lastname:  randomString(12),
		Email:     email,
		Password:  password,
	})
	req, _ := http.NewRequest("POST", "/register", bytes.NewReader(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostRegisterHandler(&repo))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var recipeUserId int
	json.NewDecoder(rr.Body).Decode(&recipeUserId)

	body, _ = json.Marshal(models.Credentials{
		Email:    email,
		Password: password,
	})

	req, _ = http.NewRequest("POST", "/login", bytes.NewReader(body))

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(PostLoginHandler(&repo))
	handler.ServeHTTP(rr, req)

	bodyBytes, _ := io.ReadAll(rr.Body)
	token := string(bodyBytes)

	if token == "" {
		t.Errorf("token should be empty, but was: %s", token)
	}

	Teardown(recipeUserId)
}
