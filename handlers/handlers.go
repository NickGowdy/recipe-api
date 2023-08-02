package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/recipe-api/repository"
)

type Handlers struct {
	userRepository   *repository.UserRepository
	recipeRepository *repository.RecipeRepository
}

func NewHandlers(userRepository *repository.UserRepository, recipeRepository *repository.RecipeRepository) *Handlers {
	return &Handlers{
		userRepository:   userRepository,
		recipeRepository: recipeRepository,
	}
}

func (r *Handlers) SetupRoutes() {
	log.Println("Loading routes...")
	mr := mux.NewRouter()

	mr.HandleFunc("/register", PostRegisterHandler(r.userRepository)).Methods("POST")
	mr.HandleFunc("/login", PostLoginHandler(r.userRepository)).Methods("POST")

	mr.Handle("/recipe", Middleware(GetAllRecipesHandler(r.recipeRepository))).Methods("GET")
	mr.Handle("/recipe/{id}", Middleware(GetRecipeHandler(r.recipeRepository))).Methods("GET")
	mr.Handle("/recipe", Middleware(InsertRecipeHandler(r.recipeRepository))).Methods("POST")
	mr.Handle("/recipe/{id}", Middleware(UpdateRecipeHandler(r.recipeRepository))).Methods("PUT")
	mr.Handle("/recipe/{id}", Middleware(DeleteRecipeHandler(r.recipeRepository))).Methods("DELETE")

	mr.HandleFunc("/health-check", HealthCheck).Methods("GET")

	http.Handle("api/", mr)
	err := http.ListenAndServe(":8080", mr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("API is now up...")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}
