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
	case len(UrlSegments) == 3 && UrlSegments[1] == "accounts" && r.Method == http.MethodGet:
		HandleAccount(w, r)
	case len(UrlSegments) == 4 && UrlSegments[1] == "accounts" && UrlSegments[3] == "recipes" && r.Method == http.MethodGet:
		HandleRecipes(w, r)
	}
}
