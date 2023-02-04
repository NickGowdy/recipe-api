package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/recipe-api/repository"
)

func SetupRoutes(repo *repository.RecipeRepository) {
	log.Println("some messaage")
	r := mux.NewRouter()

	r.HandleFunc("/register", PostRegisterHandler(repo)).Methods("POST")
	r.HandleFunc("/login", PostLoginHandler(repo)).Methods("POST")

	r.Handle("/recipe", middleware(GetRecipeHandler(repo))).Methods("GET")
	// r.HandleFunc("/recipe", GetRecipesHandler(repo)).Methods("GET")
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

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := jwt.MapClaims{}
		var jwtKey = []byte("SecretYouShouldHide")
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			tokenString := authHeader[1]
			// Parse the token
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.Write([]byte(fmt.Sprintf("Welcome %v!", claims["recipe_user_id"])))
		}
	})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}
