package main

import (
	"github.com/recipe-api/m/recipe"
)

const basePath = "/api"

func main() {
	setupRoutes()
}

func setupRoutes() {
	recipe.SetupRoutes(basePath)
}
