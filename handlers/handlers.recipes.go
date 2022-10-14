package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/recipe-api/db/repository"
	"github.com/recipe-api/models"
)

func HandleRecipes(w http.ResponseWriter, r *http.Request) {
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

func HandleRecipe(w http.ResponseWriter, r *http.Request) {
	var nr models.Recipe

	err := json.NewDecoder(r.Body).Decode(&nr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, _ := repository.SaveRecipe(&nr)

	if b {
		w.Write([]byte("works"))
	}
}
