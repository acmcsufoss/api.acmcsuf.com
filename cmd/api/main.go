package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/sqlite"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling for graceful shutdown.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		log.Println("Shutting down the server...")
		cancel()
	}()

	// Set up the database connection.
	uri, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatal("DATABASE_URL must be set")
	}

	d, err := sql.Open("sqlite", uri)
	if err != nil {
		log.Fatalf("Error opening SQLite database: %v", err)
	}

	// if err := sqlite.Migrate(ctx, db); err != nil {
	// 	return nil, errors.Wrap(err, "cannot migrate sqlite db")
	// }
	//
	// return sqliteStore{
	// 	q:   sqlite.New(db),
	// 	db:  db,
	// 	ctx: ctx,
	// }, nil

	q := sqlite.New(d)
	if err != nil {
		log.Fatalf("Error creating SQLite store: %v", err)
	}
	defer func() {
		if err := d.Close(); err != nil {
			log.Fatalf("Error closing SQLite store: %v", err)
		}
	}()

	// Initialize and start the HTTP server.
	handler := api.New(q)
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	serverAddr := fmt.Sprintf(":%s", port)
	go func() {
		fmt.Printf("Server started on http://127.0.0.1%s\n", serverAddr)
		if err := http.ListenAndServe(serverAddr, handler); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Wait for shutdown signal.
	<-ctx.Done()

	log.Println("Server shut down.")
}
