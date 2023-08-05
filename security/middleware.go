package security

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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

func GenerateToken(id int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":            json.Number(strconv.FormatInt(time.Now().Add(time.Hour*time.Duration(1)).Unix(), 10)),
		"iat":            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
		"recipe_user_id": id,
	})

	tokenString, err := token.SignedString([]byte("SecretYouShouldHide"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
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
