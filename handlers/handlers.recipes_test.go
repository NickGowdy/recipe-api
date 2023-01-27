package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
	recipeId := setupFixture()

	var recipe repository.Recipe
	db := recipeDb.NewRecipeDb()

	req, err := http.NewRequest("GET", fmt.Sprintf("/recipe/%v", recipeId), nil)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": fmt.Sprint(recipeId),
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

	if recipe.Id != recipeId {
		t.Errorf("recipe id is wrong value: got %v want %v",
			recipe.Id, 1)
	}

	if recipe.RecipeName != "Nick's recipe" {
		t.Errorf("recipe id is wrong value: got %v want %v",
			recipe.RecipeName, "Nick's recipe")
	}

	if recipe.RecipeSteps != "Some steps for Nick's recipe" {
		t.Errorf("recipe id is wrong value: got %v want %v",
			recipe.RecipeSteps, "Some steps for Nick's recipe")
	}
}

func TestInsertRecipe(t *testing.T) {
	setupEnv()

	recipeToInsert := repository.Recipe{
		Id:          0,
		AccountId:   1,
		RecipeName:  "Nick's other recipe",
		RecipeSteps: "Some other steps for Nick's recipe",
	}
	body, _ := json.Marshal(recipeToInsert)
	db := recipeDb.NewRecipeDb()

	req, err := http.NewRequest("POST", "/recipe", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(InsertRecipeHandler(db))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	req, err = http.NewRequest("GET", "/recipe/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	var id int64

	json.NewDecoder(rr.Body).Decode(&id)

	vars := map[string]string{
		"id": fmt.Sprint(id),
	}

	req = mux.SetURLVars(req, vars)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetRecipeHandler(db))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var recipe repository.Recipe
	json.NewDecoder(rr.Body).Decode(&recipe)

	if recipe.Id == 0 {
		t.Errorf("recipe id is wrong value: got %v want %v",
			0, recipe.Id)
	}

	if recipe.RecipeName != recipeToInsert.RecipeName {
		t.Errorf("recipe id is wrong value: got %v want %v",
			recipe.RecipeName, "Nick's recipe")
	}

	if recipe.RecipeSteps != recipeToInsert.RecipeSteps {
		t.Errorf("recipe id is wrong value: got %v want %v",
			recipe.RecipeSteps, "Some steps for Nick's recipe")
	}

}

func setupEnv() {
	os.Setenv("user", "postgres")
	os.Setenv("password", "postgres")
	os.Setenv("dbname", "recipes_db")
	os.Setenv("host", "localhost")
	os.Setenv("port", "5432")
}

func setupFixture() int64 {
	recipeToInsert := repository.Recipe{
		Id:          0,
		AccountId:   1,
		RecipeName:  "Nick's recipe",
		RecipeSteps: "Some steps for Nick's recipe",
	}
	body, _ := json.Marshal(recipeToInsert)
	db := recipeDb.NewRecipeDb()

	req, err := http.NewRequest("POST", "/recipe", bytes.NewReader(body))
	if err != nil {
		log.Panic(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(InsertRecipeHandler(db))

	handler.ServeHTTP(rr, req)

	var id int64

	json.NewDecoder(rr.Body).Decode(&id)
	return id
}
