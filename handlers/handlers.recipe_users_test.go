package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/recipe-api/models"
	"github.com/recipe-api/recipeDb"
	"github.com/recipe-api/repository"
)

func TestRegister(t *testing.T) {
	SetupEnv()

	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)

	register := models.Register{
		Firstname: "Test",
		Lastname:  "User",
		Email:     "testuser@gmail.com",
		Password:  "password",
	}
	body, _ := json.Marshal(register)
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
