package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/routes"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		log.Println("Shutting down the server...")
		cancel()
	}()

	// Setup SQLite database & make sure we can connect to it
	uri := os.Getenv("DATABASE_URL")
	if uri == "" {
		log.Fatal("DATABASE_URL must be set")
	}
	db, err := sql.Open("sqlite", uri)
	if err != nil {
		log.Fatalf("Error opening SQLite database: %vl", err)
	}
	defer db.Close()
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	schemaBytes, err := os.ReadFile("internal/db/sql/schemas/schema.sql")
	if err != nil {
		log.Fatalf("Error reading schema file: %v", err)
	}

	if _, err := db.ExecContext(ctx, string(schemaBytes)); err != nil {
		log.Fatalf("Error initializing db schema: %v", err)
	}

	// Now we init services & gin router, and then start the server
	// Should this be moved to the routes module??
	queries := models.New(db)
	eventsService := services.NewEventsService(queries)
	router := gin.Default()
	router.SetTrustedProxies([]string{
		"127.0.0.1/32",
	})
	routes.SetupRoutes(router, eventsService)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serverAddr := fmt.Sprintf(":%s", port)
	go func() {
		log.Printf("Server startd on http://127.0.0.1%s\n", serverAddr)
		if err := router.Run(serverAddr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Server shut down.")
}
