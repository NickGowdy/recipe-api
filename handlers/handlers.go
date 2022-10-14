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

const accountsPath = "accounts"

func SetupRoutes(apiBasePath string) {
	accountHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if strings.Contains(r.URL.Path, "/recipes") {
				handleRecipes(w, r)
			} else {
				handleAccount(w, r)
			}
		}
	})
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, accountsPath), accountHandler)
}

func handleRecipes(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "/")

	if len(urlPathSegments) == 5 {
		accountId, err := strconv.Atoi(urlPathSegments[3])

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		rs, err := repository.GetRecipes(accountId)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		j, err := json.Marshal(rs)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(j)
	}
}

func handleAccount(w http.ResponseWriter, r *http.Request) {
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
