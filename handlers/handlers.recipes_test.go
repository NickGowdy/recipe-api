package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/recipe-api/recipeDb"
	"github.com/recipe-api/recipeDb/repository"
)

func TestGetRecipe(t *testing.T) {
	setupEnv()

	var recipe repository.Recipe
	db := recipeDb.NewRecipeDb()

	req, err := http.NewRequest("GET", "/recipe/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetRecipeHandler(db))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	json.NewDecoder(rr.Body).Decode(&recipe)

	if recipe.Id != 1 {
		t.Errorf("recipe id is wrong value: got %v want %v",
			recipe.Id, 1)
	}
}

func setupEnv() {
	os.Setenv("user", "postgres")
	os.Setenv("password", "postgres")
	os.Setenv("dbname", "recipes_db")
	os.Setenv("host", "localhost")
	os.Setenv("port", "5432")
}
