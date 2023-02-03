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
		creds, shouldReturn := getCredentials(r, w)
		if shouldReturn {
			return
		}

		m, err := repo.GetRecipeUser(creds.Email, creds.Password)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(&m)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(j)

	}
	return http.HandlerFunc(fn)
}

func getCredentials(r *http.Request, w http.ResponseWriter) (models.Credentials, bool) {
	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return models.Credentials{}, true
	}
	return creds, false
}

func PostRegisterHandler(repo *repository.RecipeRepository) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var register models.Register
		if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		m, err := repo.InsertRecipeUser(&register)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(&m)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(j)
	}
	return http.HandlerFunc(fn)
}
