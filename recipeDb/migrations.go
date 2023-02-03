package recipeDb

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/qustavo/dotsql"
)

func Migrate() {
	db := NewRecipeDb()
	dotSql := getDirectory()

	fmt.Println(os.Getenv("APP_ENV"))
	fmt.Println("Running migrations")
	db.runScript(dotSql, "create-recipe_user-table")
	db.runScript(dotSql, "create-recipe-table")
	db.runScript(dotSql, "create-ingredient-table")
	db.runScript(dotSql, "create-quantity_type-table")
	db.runScript(dotSql, "create-ingredient_quantity_type-table")

	// close database
	defer db.SqlDb.Close()
}

func getDirectory() *dotsql.DotSql {
	// get relative path with runtime.caller
	_, b, _, _ := runtime.Caller(0)
	relativePath := path.Join(path.Dir(b))

	dot, err := dotsql.LoadFromFile(fmt.Sprintf("%s/init.sql", relativePath))

	if err != nil {
		log.Fatal(err)
	}

	return dot
}

func (db *RecipeDb) runScript(dot *dotsql.DotSql, name string) {
	_, err := dot.Exec(db.SqlDb, name)
	if err != nil {
		log.Fatal(err)
	}
}
