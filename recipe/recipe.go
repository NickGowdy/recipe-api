package recipe

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/recipe-api/repository"
)

type Recipe struct {
	repo *repository.RecipeRepository
}

func NewRecipe(repo repository.RecipeRepository) *Recipe {
	return &Recipe{
		repo: &repo,
	}
}

func (r Recipe) Get() http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		recipeUserId, shouldReturn := getRecipeUserId(req, w)
		if shouldReturn {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		recipeId, err := strconv.Atoi(mux.Vars(req)["id"])

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		rs, err := r.repo.GetRecipe(recipeId, recipeUserId)

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

func (r Recipe) GetAll() http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		recipeUserId, shouldReturn := getRecipeUserId(req, w)
		if shouldReturn {
			return
		}

		rs, err := r.repo.GetRecipes(recipeUserId)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
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

func (r Recipe) Insert() http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		recipeUserId, shouldReturn := getRecipeUserId(req, w)
		if shouldReturn {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var recipeToSave repository.SaveRecipe
		if err := json.NewDecoder(req.Body).Decode(&recipeToSave); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := r.repo.InsertRecipe(recipeUserId, &recipeToSave)
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

func (r Recipe) Update() http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		recipeUserId, shouldReturn := getRecipeUserId(req, w)
		if shouldReturn {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		recipeId, err := strconv.Atoi(mux.Vars(req)["id"])

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		var recipeToUpdate repository.SaveRecipe
		if err := json.NewDecoder(req.Body).Decode(&recipeToUpdate); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = r.repo.UpdateRecipe(recipeId, recipeUserId, &recipeToUpdate)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func (r Recipe) Delete() http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		recipeUserId, shouldReturn := getRecipeUserId(req, w)
		if shouldReturn {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		recipeId, err := strconv.Atoi(mux.Vars(req)["id"])

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		_, err = r.repo.DeleteRecipe(recipeId, recipeUserId)

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
	recipeUserIdFloat, ok := props["recipe_user_id"].(float64)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return 0, true
	}

	recipeUserId := int(recipeUserIdFloat)
	return recipeUserId, false
}
