package recipeDb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
)

func Migrate() {
	godotenv.Load()

	// OR: Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	datasourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, err := sql.Open("postgres", datasourceName)
	if err != nil {
		log.Fatal(err)
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
