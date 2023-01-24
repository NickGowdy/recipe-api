package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/recipe-api/db"
)

func TestGetRecipes(t *testing.T) {
	t.Setenv("user", "postgres")
	t.Setenv("password", "postgres")
	t.Setenv("dbname", "recipes_db")
	t.Setenv("host", "localhost")
	t.Setenv("port", "5432")
	db.Migrate()

	handler := http.HandlerFunc(GetRecipes)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/recipes", nil)

	if err != nil {
		t.Fatal(err)
	}

	// Then: call handler.ServeHTTP(rr, req) like in our first example.
	handler.ServeHTTP(rr, req)

	if rr.Result().StatusCode != 200 {
		t.Errorf("StatusCode should be 200 but is: %v", rr.Result().StatusCode)
	}
}
