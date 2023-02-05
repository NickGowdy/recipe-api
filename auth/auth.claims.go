package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func GetClaimsFromToken(r *http.Request) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	var jwtKey = []byte("SecretYouShouldHide")
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(authHeader) != 2 {
		return nil, fmt.Errorf("http.StatusUnauthorized")
		// w.WriteHeader(http.StatusUnauthorized)
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
