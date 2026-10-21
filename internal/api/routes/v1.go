// This file (v1.go) initializes v1 routes used by the server. Called by server.go

package routes

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/handlers"
	mw "github.com/acmcsufoss/api.acmcsuf.com/internal/api/middleware"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
)

func SetupV1(router *gin.Engine, eventService services.EventsServicer,
	boardService services.BoardServicer) {

	cfg := config.Load()
	if cfg.DiscordBotToken == "" && cfg.Env != "development" {
		log.Fatal("Error: DISCORD_BOT_TOKEN is not set")
	}

	var botSession *discordgo.Session
	var err error
	if cfg.DiscordBotToken != "" {
		botSession, err = discordgo.New("Bot " + cfg.DiscordBotToken)
		if err != nil {
			log.Fatalf("%v", err)
		}
		err = botSession.Open()
		if err != nil {
			log.Fatalf("Failed to open bot session: %v", err)
		}
		defer botSession.Close()
	}

	eh := handlers.NewEventHandler(eventService)
	bh := handlers.NewBoardHandler(boardService)

	v1 := router.Group("/v1")
	// Public (read-only) routes
	{
		v1.GET("/events", eh.GetEvents)
		v1.GET("/events/:id", eh.GetEvent)

		board := v1.Group("/board")
		{
			board.GET("/officers", bh.GetOfficers)
			board.GET("/officers/:id", bh.GetOfficer)

			board.GET("/tiers", bh.GetTiers)
			board.GET("/tiers/:tier", bh.GetTier)

			board.GET("/positions", bh.GetPositions)
			board.GET("/positions/:id", bh.GetPosition)
		}
	}

	// Protected (write) routes
	protected := v1.Group("/")
	protected.Use(mw.DiscordAuth(botSession, "Board"))
	{
		protected.POST("/events", eh.CreateEvent)
		protected.PUT("/events/:id", eh.UpdateEvent)
		protected.DELETE("/events/:id", eh.DeleteEvent)

		board := protected.Group("/board")
		{
			// Officers
			board.POST("/officers", bh.CreateOfficer)
			board.PUT("/officers/:id", bh.UpdateOfficer)
			board.DELETE("/officers/:id", bh.DeleteOfficer)

			// Tiers
			board.POST("/tiers", bh.CreateTier)
			board.PUT("/tiers/:tier", bh.UpdateTier)
			board.DELETE("/tiers/:tier", bh.DeleteTier)

			// Positions
			board.POST("/positions", bh.CreatePosition)
			board.PUT("/positions", bh.UpdatePosition)
			board.DELETE("/positions", bh.DeletePosition)
		}
	}
}
