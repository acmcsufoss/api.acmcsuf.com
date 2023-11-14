package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/acmcsufoss/api.acmcsuf.com/api/server"
	"github.com/acmcsufoss/api.acmcsuf.com/stores"
)

func main() {
	ctx := context.Background()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutting down the server...")
		os.Exit(1)
	}()

	s, err := stores.NewSQLite(ctx, "./db.sqlite") // ":memory:")
	if err != nil {
		log.Fatalf("Error creating SQLite store: %v", err)
	}

	service := server.NewOpenAPI(s)
	go func() {
		fmt.Println("Server started on http://127.0.0.1:8080")
		if err := http.ListenAndServe(":8080", service); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	if err := s.Close(); err != nil {
		log.Fatalf("Error closing SQLite store: %v", err)
	}

	log.Println("Server shut down.")

}
