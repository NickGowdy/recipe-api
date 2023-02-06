package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/recipe-api/auth"
)

var claimsKey string = "claims"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := auth.GetClaimsFromToken(r)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusUnauthorized)
		}

		ctx := context.WithValue(r.Context(), claimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
