package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/recipe-api/models"
	"github.com/recipe-api/repository"
	"golang.org/x/crypto/bcrypt"
)

func PostLoginHandler(repo *repository.RecipeRepository) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		creds, shouldReturn := getCredentials(r, w)
		if shouldReturn {
			return
		}

		ru, err := repo.GetRecipeUserPwd(creds.Email)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(ru.Password), []byte(creds.Password)); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}

		j, err := json.Marshal(&ru)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Println(j)

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

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), 8)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		hashedPasswordStr := string(hashedPassword)
		m, err := repo.InsertRecipeUser(&register.Firstname, &register.Lastname, &register.Email, &hashedPasswordStr)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusUnprocessableEntity)
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
