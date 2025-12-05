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

	// Version 1 routes
	v1 := router.Group("/v1")
	// All v1 routes are protected for now
	// TODO: We don't actually want ALL routes to be protected. Making routes like
	// /v1/events public will be much more useful for display on the website
	v1.Use(middleware.DiscordAuth(botSession, "board"))
	{
		events := v1.Group("/events")
		{
			h := handlers.NewEventHandler(eventService)
			events.GET("", h.GetEvents)
			events.GET(":id", h.GetEvent)
			events.POST("", h.CreateEvent)
			events.PUT(":id", h.UpdateEvent)
			events.DELETE(":id", h.DeleteEvent)
		}

		announcements := v1.Group("/announcements")
		{
			h := handlers.NewAnnouncementHandler(announcementService)
			announcements.GET("", h.GetAnnouncements)
			announcements.GET(":id", h.GetAnnouncement)
			announcements.POST("", h.CreateAnnouncement)
			announcements.PUT(":id", h.UpdateAnnouncement)
			announcements.DELETE(":id", h.DeleteAnnouncement)
		}

		board := v1.Group("/board")
		{
			h := handlers.NewBoardHandler(boardService)

			// Officers
			board.GET("/officers", h.GetOfficers)
			board.GET("/officers/:id", h.GetOfficer)
			board.POST("/officers", h.CreateOfficer)
			board.PUT("/officers/:id", h.UpdateOfficer)
			board.DELETE("/officers/:id", h.DeleteOfficer)

			// Tiers
			board.GET("/tiers", h.GetTiers)
			board.GET("/tiers/:id", h.GetTier)
			board.POST("/tiers", h.CreateTier)
			board.PUT("/tiers/:id", h.UpdateTier)
			board.DELETE("/tiers/:id", h.DeleteTier)

			// Positions
			board.GET("/positions", h.GetPositions)
			board.GET("/positions/:id", h.GetPosition)
			board.POST("/positions", h.CreatePosition)
			board.PUT("/positions", h.UpdatePosition)
			board.DELETE("/positions", h.DeletePosition)
		}
	}
}
