package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/recipe-api/models"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter here.....")
	switch r.Method {
	case http.MethodPost:

		var u models.User

		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		j, err := json.Marshal(u)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Write(j)
	}
}
