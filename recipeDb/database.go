package recipeDb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type RecipeDb struct {
	SqlDb *sql.DB
}

func NewRecipeDb() *RecipeDb {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Panic(err)
	}
	return &RecipeDb{SqlDb: db}
}
