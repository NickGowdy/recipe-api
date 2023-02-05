package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/recipe-api/models"
	"github.com/recipe-api/repository"
)

func GetAllRecipesHandler(repo *repository.RecipeRepository) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		recipeUserId, shouldReturn := getRecipeUserId(r, w)
		if shouldReturn {
			return
		}

		rs, err := repo.GetRecipes(recipeUserId)

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

func GetRecipeHandler(repo *repository.RecipeRepository) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		recipeUserId, shouldReturn := getRecipeUserId(r, w)
		if shouldReturn {
			return
		}
		recipeId, err := strconv.Atoi(mux.Vars(r)["id"])

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		rs, err := repo.GetRecipe(recipeId, recipeUserId)

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

func InsertRecipeHandler(repo *repository.RecipeRepository) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		recipeUserId, shouldReturn := getRecipeUserId(r, w)
		if shouldReturn {
			return
		}

		var recipeToSave models.InsertRecipe
		if err := json.NewDecoder(r.Body).Decode(&recipeToSave); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := repo.InsertRecipe(recipeUserId, &recipeToSave)
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

func UpdateRecipeHandler(repo *repository.RecipeRepository) http.HandlerFunc {
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

func DeleteRecipeHandler(repo *repository.RecipeRepository) http.HandlerFunc {
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

func getRecipeUserId(r *http.Request, w http.ResponseWriter) (int, bool) {
	props, _ := r.Context().Value("claims").(jwt.MapClaims)
	log.Print(props)
	recipeUserIdFloat, ok := props["recipe_user_id"].(float64)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return 0, true
	}

	recipeUserId := int(recipeUserIdFloat)
	return recipeUserId, false
}
