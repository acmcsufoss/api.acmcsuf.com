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
	announcementService services.AnnouncementServicer, officerService services.OfficerServicer, positionService services.PositionServicer, tierService services.TierServicer,
) {
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
	oh := handlers.NewOfficerHandler(officerService)
	ph := handlers.NewPositionHandler(positionService)
	th := handlers.NewTierHandler(tierService)

	v1 := router.Group("/v1")
	// Public (read-only) routes
	{
		v1.GET("/events", eh.GetEvents)
		v1.GET("/events/:id", eh.GetEvent)

		v1.GET("/announcements", ah.GetAnnouncements)
		v1.GET("/announcements/:id", ah.GetAnnouncement)

		board := v1.Group("/board")
		{
			board.GET("/officers", oh.GetOfficers)
			board.GET("/officers/:id", oh.GetOfficer)

			board.GET("/tiers", th.GetTiers)
			board.GET("/tiers/:id", th.GetTier)

			board.GET("/positions", ph.GetPositions)
			board.GET("/positions/:id", ph.GetPosition)
		}
	}

	// Protected (write) routes
	protected := v1.Group("/")
	protected.Use(mw.DiscordAuth(botSession, "Board"))
	{
		protected.POST("/events", eh.CreateEvent)
		protected.PUT("/events/:id", eh.UpdateEvent)
		protected.DELETE("/events/:id", eh.DeleteEvent)

		protected.POST("/announcements", ah.CreateAnnouncement)
		protected.PUT("/announcements/:id", ah.UpdateAnnouncement)
		protected.DELETE("/announcements/:id", ah.DeleteAnnouncement)

		board := protected.Group("/board")
		{
			// Officers
			board.POST("/officers", oh.CreateOfficer)
			board.PUT("/officers/:id", oh.UpdateOfficer)
			board.DELETE("/officers/:id", oh.DeleteOfficer)

			// Tiers
			board.POST("/tiers", th.CreateTier)
			board.PUT("/tiers/:id", th.UpdateTier)
			board.DELETE("/tiers/:id", th.DeleteTier)

			// Positions
			board.POST("/positions", ph.CreatePosition)
			board.PUT("/positions", ph.UpdatePosition)
			board.DELETE("/positions", ph.DeletePosition)
		}
	}
}
