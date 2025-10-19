package db

//go:generate sqlc generate

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

func New(ctx context.Context) (*sql.DB, func(), error) {
	uri := os.Getenv("DATABASE_URL")
	if uri == "" {
		return nil, nil, errors.New("DATABASE_URL must be set")
	}

	db, err := sql.Open("sqlite", uri)
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
