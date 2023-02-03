package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/recipe-api/models"
	"github.com/recipe-api/repository"
)

func PostLoginHandler(repo *repository.RecipeRepository) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var user models.RecipeUser
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err := repo.GetUser(user.Email, user.Password)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(make([]byte, 0))

	}
	return http.HandlerFunc(fn)
}
