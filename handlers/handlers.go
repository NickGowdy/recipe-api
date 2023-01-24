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
	r.HandleFunc("/recipe", GetRecipes).Methods("GET")
	r.HandleFunc("/recipe/{id}", GetRecipe).Methods("GET")
	r.HandleFunc("/recipe", InsertRecipe).Methods("POST")
	r.HandleFunc("/recipe/{id}", UpdateRecipe).Methods("PUT")
	r.HandleFunc("/recipe/{id}", DeleteRecipe).Methods("DELETE")
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

func GetRecipes(w http.ResponseWriter, r *http.Request) {
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
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	recipeId, err := strconv.Atoi(mux.Vars(r)["id"])

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

func InsertRecipe(w http.ResponseWriter, r *http.Request) {
	var recipeToSave repository.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipeToSave); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := repository.InsertRecipe(&recipeToSave)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	recipeId, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	var recipeToSave repository.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipeToSave); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = repository.UpdateRecipe(&recipeToSave, recipeId)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	recipeId, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err = repository.DeleteRecipe(recipeId)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
