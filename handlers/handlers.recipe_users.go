package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/recipe-api/models"
	"github.com/recipe-api/repository"
	"golang.org/x/crypto/bcrypt"
)

func PostRegisterHandler(repo *repository.UserRepository) http.HandlerFunc {
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
		m, err := repo.InsertRecipeUser(register.Firstname, register.Lastname, register.Email, hashedPasswordStr)

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

		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
	return http.HandlerFunc(fn)
}

func PostLoginHandler(repo *repository.UserRepository) http.HandlerFunc {
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

		tokenString, err := generateJWT(int64(ru.ID))
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(tokenString))
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

func generateJWT(id int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":            json.Number(strconv.FormatInt(time.Now().Add(time.Hour*time.Duration(1)).Unix(), 10)),
		"iat":            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
		"recipe_user_id": id,
	})

	tokenString, err := token.SignedString([]byte("SecretYouShouldHide"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
