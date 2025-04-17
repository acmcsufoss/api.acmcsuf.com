package db

//go:generate sqlc generate

import (
	"context"
	"os"
	"database/sql"
	"errors"
	"fmt"

	_ "modernc.org/sqlite"
)

func New(ctx context.Context) (*sql.DB, error) {
	uri := os.Getenv("DATABASE_URL")
	if uri == "" {
		return nil, errors.New("DATABASE_URL must be set")
	}

	db, err := sql.Open("sqlite", uri)
	if err != nil {
		return nil, fmt.Errorf("Error opening SQLite database: %v", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}

	schemaBytes, err := os.ReadFile("internal/db/sql/schemas/schema.sql")
	if err != nil {
		return nil, fmt.Errorf("Error reading schema file: %v", err)
	}

	if _, err := db.ExecContext(ctx, string(schemaBytes)); err != nil {
		return nil, fmt.Errorf("Error initializing db schema: %v", err)

	}

	return db, nil
}
