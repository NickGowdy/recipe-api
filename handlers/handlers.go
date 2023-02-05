package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/recipe-api/repository"
)

func SetupRoutes(repo *repository.RecipeRepository) {
	log.Println("Loading routes...")
	r := mux.NewRouter()

	r.HandleFunc("/register", PostRegisterHandler(repo)).Methods("POST")
	r.HandleFunc("/login", PostLoginHandler(repo)).Methods("POST")

	r.Handle("/recipe", middleware(GetAllRecipesHandler(repo))).Methods("GET")
	r.Handle("/recipe/{id}", middleware(GetRecipeHandler(repo))).Methods("GET")
	r.Handle("/recipe", middleware(InsertRecipeHandler(repo))).Methods("POST")
	r.Handle("/recipe/{id}", middleware(UpdateRecipeHandler(repo))).Methods("PUT")
	r.Handle("/recipe/{id}", middleware(UpdateRecipeHandler(repo))).Methods("DELETE")

	r.HandleFunc("/health-check", HealthCheck).Methods("GET")

	http.Handle("api/", r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("API is now up...")
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := jwt.MapClaims{}
		var jwtKey = []byte("SecretYouShouldHide")
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			tokenString := authHeader[1]
			_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "claims", claims)
			// Access context values in handlers like this
			// props, _ := r.Context().Value("props").(jwt.MapClaims)
			next.ServeHTTP(w, r.WithContext(ctx))

			return
		}
	})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}
