package main

import (
	"context"
	"log"
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

	s, err := stores.NewSQLite(ctx, ":memory:")
	if err != nil {
		log.Fatalf("Error creating SQLite store: %v", err)
	}

	h := server.NewHandler(server.HandlerOptions{
		Ctx:   ctx,
		Store: s,
		Port:  ":8080",
	})

	go func() { h.Serve() }()

	<-ctx.Done()

	if err := s.Close(); err != nil {
		log.Fatalf("Error closing SQLite store: %v", err)
	}

	log.Println("Server shut down.")

}
