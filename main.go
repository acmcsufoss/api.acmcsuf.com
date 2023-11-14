package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/acmcsufoss/api.acmcsuf.com/api/openapi"
	"github.com/acmcsufoss/api.acmcsuf.com/stores"
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

	// Initialize SQLite store.
	s, err := stores.NewSQLite(ctx, "./db.sqlite") // ":memory:")
	if err != nil {
		log.Fatalf("Error creating SQLite store: %v", err)
	}
	defer func() {
		if err := s.Close(); err != nil {
			log.Fatalf("Error closing SQLite store: %v", err)
		}
	}()

	// Initialize and start the HTTP server.
	handler := openapi.NewOpenAPI(s)
	serverAddr := ":8080"
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
