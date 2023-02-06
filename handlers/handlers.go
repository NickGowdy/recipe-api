package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/recipe-api/repository"
)

func SetupRoutes(repo *repository.RecipeRepository) {
	log.Println("Loading routes...")
	r := mux.NewRouter()

	r.HandleFunc("/register", PostRegisterHandler(repo)).Methods("POST")
	r.HandleFunc("/login", PostLoginHandler(repo)).Methods("POST")

	r.Handle("/recipe", Middleware(GetAllRecipesHandler(repo))).Methods("GET")
	r.Handle("/recipe/{id}", Middleware(GetRecipeHandler(repo))).Methods("GET")
	r.Handle("/recipe", Middleware(InsertRecipeHandler(repo))).Methods("POST")
	r.Handle("/recipe/{id}", Middleware(UpdateRecipeHandler(repo))).Methods("PUT")
	r.Handle("/recipe/{id}", Middleware(DeleteRecipeHandler(repo))).Methods("DELETE")

	r.HandleFunc("/health-check", HealthCheck).Methods("GET")

	http.Handle("api/", r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("API is now up...")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}
