package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/recipe-api/db/repository"
)

func handleAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", "accounts"))
		if len(urlPathSegments[1:]) > 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		accountId, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		returnedAccount, err := repository.GetAccount(accountId)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		j, err := json.Marshal(returnedAccount)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(j)
	}
}
