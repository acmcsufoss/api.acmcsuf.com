package db

//go:generate sqlc generate

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

func New(ctx context.Context, url string) (*sql.DB, func(), error) {
	db, err := sql.Open("sqlite", url)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening SQLite database: %v", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, nil, fmt.Errorf("error connecting to database: %v", err)
	}

	schemaBytes, err := os.ReadFile("internal/db/sql/schemas/schema.sql")
	if err != nil {
		return nil, nil, fmt.Errorf("error reading schema file: %v", err)
	}

	if _, err := db.ExecContext(ctx, string(schemaBytes)); err != nil {
		return nil, nil, fmt.Errorf("error initializing db schema: %v", err)

	}

	return db, func() {
		db.Close()
	}, nil
}
