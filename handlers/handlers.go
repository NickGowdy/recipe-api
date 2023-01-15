package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func SetupRoutes(apiBasePath string) {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getHandler(w, r)
	})

	http.Handle(fmt.Sprintf("%s/", apiBasePath), handler)
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	UrlSegments := strings.Split(r.URL.Path, "/")[1:]

	switch {
	case len(UrlSegments) == 2 && UrlSegments[1] == "user":
		HandleUser(w, r)
	case len(UrlSegments) == 2 && UrlSegments[1] == "login":
		HandleLogin(w, r)
	case len(UrlSegments) == 3 && UrlSegments[1] == "accounts":
		HandleAccount(w, r)
	case len(UrlSegments) == 4 && UrlSegments[1] == "accounts" && UrlSegments[3] == "recipes":
		HandleRecipes(w, r)
	case len(UrlSegments) == 5 && UrlSegments[1] == "accounts" && UrlSegments[3] == "recipes":
		HandleRecipe(w, r)
	case len(UrlSegments) == 2 && UrlSegments[1] == "recipes":
		HandleRecipe(w, r)
	}
}
