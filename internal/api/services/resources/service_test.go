package resources

import (
	"database/sql"
	"log"
	"testing"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/db/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func Test_GetResource(t *testing.T) {
	d, err := sql.Open("sqlite3", "../../db/sqlite/api.db")
	if err != nil {
		log.Fatalf("Error opening SQLite database: %v", err)
	}
	obj := sqlite.New(d)
	serviceObj := New(obj)
	res, err := serviceObj.GetResource("1")
	if err != nil {
		t.Fatalf("%v Error bro", err.Error())
	}
	t.Log(res.CreatedAt)
	t.Log(d, "Hello")
}
