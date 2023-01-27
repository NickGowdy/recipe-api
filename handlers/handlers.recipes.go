package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/recipe-api/models"
	"github.com/recipe-api/recipeDb/repository"
)

func GetRecipesHandler(repo *repository.Repository) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		rs, err := repo.GetRecipes()

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

func GetRecipeHandler(repo *repository.Repository) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		recipeId, err := strconv.Atoi(mux.Vars(r)["id"])

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		rs, err := repo.GetRecipe(recipeId)

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

func InsertRecipeHandler(repo *repository.Repository) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var recipeToSave models.Recipe
		if err := json.NewDecoder(r.Body).Decode(&recipeToSave); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := repo.InsertRecipe(&recipeToSave)
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

func UpdateRecipeHandler(repo *repository.Repository) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		recipeId, err := strconv.Atoi(mux.Vars(r)["id"])

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		var recipeToSave models.Recipe
		if err := json.NewDecoder(r.Body).Decode(&recipeToSave); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = repo.UpdateRecipe(&recipeToSave, recipeId)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func DeleteRecipeHandler(repo *repository.Repository) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		recipeId, err := strconv.Atoi(mux.Vars(r)["id"])

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		_, err = repo.DeleteRecipe(recipeId)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}
