package main

import (
	"github.com/recipe-api/db"
	"github.com/recipe-api/handlers"
)

func main() {
	InitApp()
}

func InitApp() {
	db.Migrate()
	handlers.SetupRoutes()
}
