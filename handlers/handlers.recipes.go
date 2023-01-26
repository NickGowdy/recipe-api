package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/recipe-api/recipeDb"
	"github.com/recipe-api/recipeDb/repository"
)

func GetRecipesHandler(db *recipeDb.RecipeDb) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		rs, err := repository.GetRecipes(db)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		j, err := json.Marshal(rs)
		if err != nil {
			log.Print(err)
		}
		w.Write(j)
	}
	return http.HandlerFunc(fn)
}

func GetRecipeHandler(db *recipeDb.RecipeDb) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		recipeId, err := strconv.Atoi(mux.Vars(r)["id"])

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		rs, err := repository.GetRecipe(db, recipeId)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		j, err := json.Marshal(rs)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(j)
	}
	return http.HandlerFunc(fn)
}

func InsertRecipeHandler(db *recipeDb.RecipeDb) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var recipeToSave repository.Recipe
		if err := json.NewDecoder(r.Body).Decode(&recipeToSave); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := repository.InsertRecipe(db, &recipeToSave)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		j, err := json.Marshal(id)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(j)
	}
	return http.HandlerFunc(fn)
}

func UpdateRecipeHandler(db *recipeDb.RecipeDb) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		recipeId, err := strconv.Atoi(mux.Vars(r)["id"])

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		var recipeToSave repository.Recipe
		if err := json.NewDecoder(r.Body).Decode(&recipeToSave); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = repository.UpdateRecipe(db, &recipeToSave, recipeId)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func DeleteRecipeHandler(db *recipeDb.RecipeDb) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		recipeId, err := strconv.Atoi(mux.Vars(r)["id"])

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		_, err = repository.DeleteRecipe(db, recipeId)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}
