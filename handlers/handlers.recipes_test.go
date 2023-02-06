package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/recipe-api/models"
	"github.com/recipe-api/recipeDb"
	"github.com/recipe-api/repository"
)

func TestGetRecipe(t *testing.T) {
	SetupEnvVars()
	testUser := SetupUser()
	token := SetupToken(testUser)
	bearer := "Bearer " + token
	recipeId := SetupRecipe(bearer)
	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)

	var recipe models.Recipe

	req, err := http.NewRequest("GET", fmt.Sprintf("/recipe/%v", recipeId), nil)
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": fmt.Sprint(recipeId),
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.Handler(Middleware(GetRecipeHandler(&repo)))

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
	TeardownRecipe(recipeId)
	TeardownUser(testUser.email)
}

func TestInsertRecipe(t *testing.T) {
	SetupEnvVars()
	testUser := SetupUser()
	token := SetupToken(testUser)
	bearer := "Bearer " + token
	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)

	recipeToInsert := models.Recipe{
		RecipeName:  "Nick's other recipe",
		RecipeSteps: "Some other steps for Nick's recipe",
	}

	body, _ := json.Marshal(recipeToInsert)

	req, err := http.NewRequest("POST", "/recipe", bytes.NewReader(body))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(Middleware(InsertRecipeHandler(&repo)))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	req, err = http.NewRequest("GET", "/recipe/1", nil)
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

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
	handler = http.Handler(Middleware(GetRecipeHandler(&repo)))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var recipe models.Recipe
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
	TeardownRecipe(recipeId)
	TeardownUser(testUser.email)
}

func TestUpdateRecipe(t *testing.T) {
	SetupEnvVars()
	testUser := SetupUser()
	token := SetupToken(testUser)
	bearer := "Bearer " + token
	recipeId := SetupRecipe(bearer)
	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)

	var recipe models.Recipe

	recipeToUpdate := models.Recipe{
		Id:          recipeId,
		RecipeName:  "This is the new name",
		RecipeSteps: "These are the new steps",
	}

	body, _ := json.Marshal(recipeToUpdate)

	req, err := http.NewRequest("PUT", fmt.Sprintf("/recipe/%v", recipeId), bytes.NewReader(body))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": fmt.Sprint(recipeId),
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.Handler(Middleware(UpdateRecipeHandler(&repo)))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	req = mux.SetURLVars(req, vars)

	rr = httptest.NewRecorder()
	handler = http.Handler(Middleware(GetRecipeHandler(&repo)))

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
	TeardownRecipe(recipeId)
	TeardownUser(testUser.email)
}

func TestDeleteRecipe(t *testing.T) {
	SetupEnvVars()
	testUser := SetupUser()
	token := SetupToken(testUser)
	bearer := "Bearer " + token
	recipeId := SetupRecipe(bearer)

	db := recipeDb.NewRecipeDb()
	repo := repository.NewRecipeRepository(db)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/recipe/%v", recipeId), nil)
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": fmt.Sprint(recipeId),
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.Handler(Middleware(DeleteRecipeHandler(&repo)))

	handler.ServeHTTP(rr, req)

	if rr.Result().StatusCode != 200 {
		fmt.Printf("error expected: %v but got %v", 200, rr.Result().StatusCode)
	}

	req, err = http.NewRequest("GET", fmt.Sprintf("/recipe/%v", recipeId), nil)
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, vars)

	rr = httptest.NewRecorder()
	handler = http.Handler(Middleware(GetRecipeHandler(&repo)))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	TeardownRecipe(recipeId)
	TeardownUser(testUser.email)
}
