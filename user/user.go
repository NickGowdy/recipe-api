package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/recipe-api/db/repository"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		urlFragment := strings.Split(r.URL.Path, "users/")
		if len(urlFragment[1:]) > 1 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		userId, err := strconv.Atoi(urlFragment[len(urlFragment)-1])
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		user, err := repository.GetUser(userId)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		j, err := json.Marshal(user)

		if err != nil {
			log.Fatal(err)
		}

		w.Write(j)
	}

}
