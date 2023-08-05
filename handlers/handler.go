package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/recipe-api/recipe"
	"github.com/recipe-api/security"
	"github.com/recipe-api/user"
)

type Handler struct {
	user   *user.User
	recipe *recipe.Recipe
}

func NewHandler(user user.User, recipe recipe.Recipe) *Handler {
	return &Handler{user: &user, recipe: &recipe}
}

func (h *Handler) Start(muxRouter *mux.Router) {
	muxRouter.HandleFunc("/register", h.registerHandler).Methods("POST")
	muxRouter.HandleFunc("/login", h.loginHandler).Methods("POST")

	muxRouter.Handle("/recipe", security.VerifyToken(recipe.Get())).Methods("GET")
	muxRouter.Handle("/recipe/{id}", security.VerifyToken(recipe.GetAll())).Methods("GET")
	muxRouter.Handle("/recipe", security.VerifyToken(recipe.Insert())).Methods("POST")
	muxRouter.Handle("/recipe/{id}", security.VerifyToken(recipe.Update())).Methods("PUT")
	muxRouter.Handle("/recipe/{id}", security.VerifyToken(recipe.Delete())).Methods("DELETE")

	muxRouter.HandleFunc("/health-check", HealthCheck).Methods("GET")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}
