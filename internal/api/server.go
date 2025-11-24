// This file (server.go) contains server initialization logic that's called by main.go

package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/routes"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

// Run initializes the database, services, and router, then starts the server.
// It waits for the context to be canceled to initiate a graceful shutdown.
func Run(ctx context.Context) {
	cfg := config.Load()
	botToken := os.Getenv("DISCORD_BOT_TOKEN")

	if botToken == "" && cfg.Env != "development" {
		log.Fatal("Error: DISCORD_BOT_TOKEN is not set")
	}
	var botSession *discordgo.Session
	if botToken != "" {
		botSession, err := discordgo.New("Bot " + botToken)
		if err != nil {
			log.Fatalf("%v", err)
		}
		botSession.Open()
		defer botSession.Close()
	}

	db, closer, err := db.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	// Now we init services & gin router, and then start the server
	queries := models.New(db)
	eventsService := services.NewEventsService(queries)
	announcementService := services.NewAnnouncementService(queries)
	boardService := services.NewBoardService(queries, db)
	router := gin.Default()

	router.SetTrustedProxies(cfg.TrustedProxies)
	routes.SetupRoot(router)
	routes.SetupV1(router, eventsService, announcementService, boardService)

	go func() {
		serverAddr := fmt.Sprintf("localhost:%s", cfg.Port)
		log.Printf("\x1b[32mServer started on http://%s\x1b[0m", serverAddr)

		if err := router.Run(serverAddr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// This is a blocking call that prevents the function from finishing until the signal
	// is received.
	<-ctx.Done()
	log.Println("\x1b[32mServer shut down.\x1b[0m")
}
