package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/routes"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/gin-gonic/gin"
)

// Run initializes the database, services, and router, then starts the server.
// It waits for the context to be canceled to initiate a graceful shutdown.
func Run(ctx context.Context) {
	db, closer, err := db.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	// Now we init services & gin router, and then start the server
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
	go func() {
		serverAddr := fmt.Sprintf("localhost:%s", port)
		log.Printf("[32m Server started on http://%s [0m ", serverAddr)

		if err := router.Run(serverAddr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Server shut down.")
}
