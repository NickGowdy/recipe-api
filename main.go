package main

import (
	"net/http"

	"github.com/recipe-api/db"
	"github.com/recipe-api/handlers"
)

const basePath = "/api"

func main() {
	setupRoutes()
}

func setupRoutes() {

	db.Migrate()

	handlers.SetupRoutes(basePath)
	http.ListenAndServe(":8080", nil)
}
