package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/routes"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db"
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

	db, err := db.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Now we init services & gin router, and then start the server
	// Should this be moved to the routes module??
	queries := models.New(db)
	eventsService := services.NewEventsService(queries)
	announcementService := services.NewAnnouncementService(queries)
	router := gin.Default()
	router.SetTrustedProxies([]string{
		"127.0.0.1/32",
	})
	routes.SetupRoutes(router, eventsService, announcementService)
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
