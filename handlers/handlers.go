package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/recipe-api/recipeDb"
)

func SetupRoutes(db *recipeDb.RecipeDb) {
	log.Println("some messaage")
	r := mux.NewRouter()
	r.HandleFunc("/recipe", GetRecipesHandler(db)).Methods("GET")
	r.HandleFunc("/recipe/{id}", GetRecipeHandler(db)).Methods("GET")
	r.HandleFunc("/recipe", InsertRecipeHandler(db)).Methods("POST")
	r.HandleFunc("/recipe/{id}", UpdateRecipeHandler(db)).Methods("PUT")
	r.HandleFunc("/recipe/{id}", DeleteRecipeHandler(db)).Methods("DELETE")

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
