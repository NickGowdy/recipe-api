package main

import (
	"net/http"

	"github.com/recipe-api/m/db"
	"github.com/recipe-api/m/handlers"
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
