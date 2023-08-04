package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var claimsKey string = "claims"

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := GetClaimsFromToken(r)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusUnauthorized)
		}

		ctx := context.WithValue(r.Context(), claimsKey, claims)
		// Access context values in handlers like this
		// props, _ := r.Context().Value("props").(jwt.MapClaims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetClaimsFromToken(r *http.Request) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	var jwtKey = []byte("SecretYouShouldHide")
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(authHeader) != 2 {
		return nil, fmt.Errorf("http.StatusUnauthorized")
	} else {
		tokenString := authHeader[1]
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			return nil, err
		}

		return claims, nil
	}
}
