// This file (v1.go) initializes v1 routes used by the server. Called by server.go

package routes

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/handlers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/middleware"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
)

func SetupV1(router *gin.Engine, eventService services.EventsServicer,
	announcementService services.AnnouncementServicer, boardService services.BoardServicer) {

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
	ah := handlers.NewAnnouncementHandler(announcementService)
	bh := handlers.NewBoardHandler(boardService)

	// Public version 1 routes (read-only stuff)
	publicV1 := router.Group("/v1")
	{
		publicV1.GET("/events", eh.GetEvents)
		publicV1.GET("/events/:id", eh.GetEvent)

		publicV1.GET("/announcements", ah.GetAnnouncements)
		publicV1.GET("/announcements/:id", ah.GetAnnouncements)

		board := publicV1.Group("/board")
		{
			board.GET("/officers", bh.GetOfficers)
			board.GET("/officers/:id", bh.GetOfficer)

			board.GET("/tiers", bh.GetTiers)
			board.GET("/tiers/:id", bh.GetTier)

			board.GET("/positions", bh.GetPositions)
			board.GET("/positions/:id", bh.GetPosition)
		}
	}

	// Protected version 1 routes (write operations)
	protectedV1 := router.Group("/v1")
	protectedV1.Use(middleware.DiscordAuth(botSession, "Board"))
	{
		protectedV1.POST("/events", eh.CreateEvent)
		protectedV1.PUT("/events/:id", eh.UpdateEvent)
		protectedV1.DELETE("/events/:id", eh.DeleteEvent)

		protectedV1.POST("announcements", ah.CreateAnnouncement)
		protectedV1.PUT("announcements/:id", ah.UpdateAnnouncement)
		protectedV1.DELETE("announcements/:id", ah.DeleteAnnouncement)

		board := protectedV1.Group("/board")
		{
			// Officers
			board.POST("/officers", bh.CreateOfficer)
			board.PUT("/officers/:id", bh.UpdateOfficer)
			board.DELETE("/officers/:id", bh.DeleteOfficer)

			// Tiers
			board.POST("/tiers", bh.CreateTier)
			board.PUT("/tiers/:id", bh.UpdateTier)
			board.DELETE("/tiers/:id", bh.DeleteTier)

			// Positions
			board.POST("/positions", bh.CreatePosition)
			board.PUT("/positions", bh.UpdatePosition)
			board.DELETE("/positions", bh.DeletePosition)
		}

	}
}
