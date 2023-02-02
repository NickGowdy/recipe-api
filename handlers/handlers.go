package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/recipe-api/repository"
)

func SetupRoutes(repo *repository.RecipeRepository) {
	log.Println("some messaage")
	r := mux.NewRouter()

	r.HandleFunc("/login", PostLoginHandler(repo)).Methods("POST")

	r.HandleFunc("/recipe", GetRecipesHandler(repo)).Methods("GET")
	r.HandleFunc("/recipe/{id}", GetRecipeHandler(repo)).Methods("GET")
	r.HandleFunc("/recipe", InsertRecipeHandler(repo)).Methods("POST")
	r.HandleFunc("/recipe/{id}", UpdateRecipeHandler(repo)).Methods("PUT")
	r.HandleFunc("/recipe/{id}", DeleteRecipeHandler(repo)).Methods("DELETE")

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
