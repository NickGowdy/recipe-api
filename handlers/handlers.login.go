package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter here.....")
	switch r.Method {
	case http.MethodPost:

		var u user

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
