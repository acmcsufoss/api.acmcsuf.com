package sqlite

import (
	"context"
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func Test_DBQueries(t *testing.T) {
	d, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Fatalf("Error opening SQLite database: %v", err)
	}
	obj := New(d)
	res, err := obj.GetResource(context.TODO(), "1")
	if err != nil {
		t.Fatalf("%v Error bro", err.Error())
	}
	t.Log(res.CreatedAt)
	t.Log(d, "Hello")
}
