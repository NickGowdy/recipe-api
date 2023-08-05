package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/recipe-api/repository"
)

func (h Handler) registerHandler(w http.ResponseWriter, r *http.Request) {
	var register repository.Register
	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)

	}

	if register.Firstname == "" {
		w.Write([]byte("firstname is a required field"))
		w.WriteHeader(http.StatusBadRequest)
	}

	if register.Lastname == "" {
		w.Write([]byte("lastname is a required field"))
		w.WriteHeader(http.StatusBadRequest)
	}

	if register.Email == "" {
		w.Write([]byte("email is a required field"))
		w.WriteHeader(http.StatusBadRequest)
	}

	if register.Password == "" {
		w.Write([]byte("password is a required field"))
		w.WriteHeader(http.StatusBadRequest)
	}

	userId, err := h.user.Register(register.Firstname, register.Lastname, register.Email, register.Password)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	j, err := json.Marshal(&userId)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func (h Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	creds, shouldReturn := getCredentials(r, w)
	if shouldReturn {
		return
	}

	token, err := h.user.Login(creds.Email, creds.Password)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write([]byte(*token))
}

func getCredentials(r *http.Request, w http.ResponseWriter) (repository.Credentials, bool) {
	var creds repository.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return repository.Credentials{}, true
	}
	return creds, false
}
