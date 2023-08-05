package recipe

import (
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/recipe-api/database"
	"github.com/recipe-api/repository"
)

type Recipe struct {
	repo *repository.RecipeRepository
}

func NewRecipe(repo *repository.RecipeRepository) *Recipe {
	return &Recipe{
		repo: repo,
	}
}

func (r Recipe) Get(recipeId int, recipeUserId int) (*database.Recipe, error) {
	recipe, err := r.repo.GetRecipe(recipeId, recipeUserId)

	if err != nil {
		return nil, err
	}
	return recipe, err
}

func (r Recipe) GetAll(recipeUserId int) (*[]database.Recipe, error) {
	recipes, err := r.repo.GetRecipes(recipeUserId)

	if err != nil {
		return nil, err
	}
	return &recipes, err
}

func (r Recipe) Insert(recipeUserId int, saveRecipe repository.SaveRecipe) (int32, error) {
	id, err := r.repo.InsertRecipe(recipeUserId, &saveRecipe)

	if err != nil {
		return 0, err
	}
	return id, err
}

func (r Recipe) Update(recipeUserId int, recipeId int, saveRecipe repository.SaveRecipe) (bool, error) {
	_, err := r.repo.UpdateRecipe(recipeId, recipeUserId, &saveRecipe)

	if err != nil {
		return false, err
	}
	return true, nil
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
