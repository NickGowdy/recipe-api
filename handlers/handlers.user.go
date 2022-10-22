package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/recipe-api/models"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var u models.User

		json.NewDecoder(r.Body).Decode(&u)

		j, err := json.Marshal(u)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Write(j)
	}
}
