package main

import (
	"github.com/recipe-api/m/db"
	"github.com/recipe-api/m/recipe"
)

const basePath = "/api"

func main() {
	setupRoutes()
}

func setupRoutes() {
	db.Migrate()

	recipe.SetupRoutes(basePath)
}
