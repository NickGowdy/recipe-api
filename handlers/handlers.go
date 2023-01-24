package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/recipe-api/db/repository"
)

func SetupRoutes() {
	log.Println("some messaage")
	r := mux.NewRouter()
	r.HandleFunc("/recipe", HandleRecipes).Methods("GET", "POST")
	r.HandleFunc("/recipe/{id}", HandleRecipe).Methods("GET")
	r.HandleFunc("/health-check", HealthCheck).Methods("GET")
	http.Handle("api/", r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func HandleRecipes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rs, err := repository.GetRecipes()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		j, err := json.Marshal(rs)
		if err != nil {
			log.Print(err)
		}
		w.Write(j)
	case "POST":
		var recipeToSave repository.Recipe
		if err := json.NewDecoder(r.Body).Decode(&recipeToSave); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		r, err := repository.InsertRecipe(&recipeToSave)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		j, err := json.Marshal(r)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func HandleRecipe(w http.ResponseWriter, r *http.Request) {
	recipeId, err := strconv.Atoi(mux.Vars(r)["id"])

	log.Print(recipeId)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	rs, err := repository.GetRecipe(recipeId)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	j, err := json.Marshal(rs)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(j)
}
