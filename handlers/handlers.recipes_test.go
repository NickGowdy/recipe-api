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

	teardownFixture(recipeId)
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

	var recipeId int64

	json.NewDecoder(rr.Body).Decode(&recipeId)

	vars := map[string]string{
		"id": fmt.Sprint(recipeId),
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
			recipe.RecipeName, recipeToInsert.RecipeName)
	}

	if recipe.RecipeSteps != recipeToInsert.RecipeSteps {
		t.Errorf("recipe id is wrong value: got %v want %v",
			recipe.RecipeSteps, recipeToInsert.RecipeSteps)
	}

	teardownFixture(recipeId)
}

func TestUpdateRecipe(t *testing.T) {
	setupEnv()
	recipeId := setupFixture()

	var recipe repository.Recipe
	db := recipeDb.NewRecipeDb()

	recipeToUpdate := repository.Recipe{
		Id:          recipeId,
		AccountId:   1,
		RecipeName:  "This is the new name",
		RecipeSteps: "These are the new steps",
	}

	body, _ := json.Marshal(recipeToUpdate)

	req, err := http.NewRequest("PUT", fmt.Sprintf("/recipe/%v", recipeId), bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": fmt.Sprint(recipeId),
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateRecipeHandler(db))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	req = mux.SetURLVars(req, vars)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetRecipeHandler(db))

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

	if recipe.RecipeName != recipeToUpdate.RecipeName {
		t.Errorf("recipe id is wrong value: got %v want %v",
			recipe.RecipeName, recipeToUpdate.RecipeName)
	}

	if recipe.RecipeSteps != recipeToUpdate.RecipeSteps {
		t.Errorf("recipe id is wrong value: got %v want %v",
			recipe.RecipeSteps, recipeToUpdate.RecipeSteps)
	}

	teardownFixture(recipeId)
}

func TestDeleteRecipe(t *testing.T) {
	setupEnv()
	recipeId := setupFixture()

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/recipe/%v", recipeId), nil)
	db := recipeDb.NewRecipeDb()
	if err != nil {
		log.Panic(err)
	}

	vars := map[string]string{
		"id": fmt.Sprint(recipeId),
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteRecipeHandler(db))

	handler.ServeHTTP(rr, req)

	if rr.Result().StatusCode != 200 {
		fmt.Printf("error expected: %v but got %v", 200, rr.Result().StatusCode)
	}

	req, err = http.NewRequest("GET", fmt.Sprintf("/recipe/%v", recipeId), nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, vars)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetRecipeHandler(db))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
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

func teardownFixture(recipeId int64) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/recipe/%v", recipeId), nil)
	db := recipeDb.NewRecipeDb()
	if err != nil {
		log.Panic(err)
	}

	vars := map[string]string{
		"id": fmt.Sprint(recipeId),
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteRecipeHandler(db))

	handler.ServeHTTP(rr, req)

	if rr.Result().StatusCode != 200 {
		fmt.Printf("error with teardown fixture expected: %v but got %v", 200, rr.Result().StatusCode)
	}
}
