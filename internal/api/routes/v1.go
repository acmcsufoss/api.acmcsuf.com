// This file (v1.go) initializes v1 routes used by the server. Called by server.go

package routes

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/handlers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/gin-gonic/gin"
)

func SetupV1Routes(router *gin.Engine, eventService services.EventsServicer,
	announcementService services.AnnouncementServicer, boardService services.BoardServicer) {
	router.GET("/swagger/*any", handlers.NewSwaggerHandler())

	// Version 1 routes
	v1 := router.Group("/v1")
	{
		eventHandler := handlers.NewEventHandler(eventService)
		v1.GET("/events", eventHandler.GetEvents)
		v1.GET("/events/:id", eventHandler.GetEvent)
		v1.POST("/events", eventHandler.CreateEvent)
		v1.PUT("/events/:id", eventHandler.UpdateEvent)
		v1.DELETE("/events/:id", eventHandler.DeleteEvent)

		announcementHandler := handlers.NewAnnouncementHandler(announcementService)
		v1.GET("/announcements", announcementHandler.GetAnnouncements)
		v1.GET("/announcements/:id", announcementHandler.GetAnnouncement)
		v1.POST("/announcements", announcementHandler.CreateAnnouncement)
		v1.PUT("/announcements/:id", announcementHandler.UpdateAnnouncement)
		v1.DELETE("/announcements/:id", announcementHandler.DeleteAnnouncement)

		boardHandler := handlers.NewBoardHandler(boardService)
		router.GET("/officers", boardHandler.GetOfficers)
		router.GET("/officer/:id", boardHandler.GetOfficer)
		router.POST("/officer", boardHandler.CreateOfficer)
		router.PUT("/officer", boardHandler.UpdateOfficer)
		router.DELETE("/officer", boardHandler.DeleteOfficer)

		router.GET("/tier/:id", boardHandler.GetTier)
		router.POST("/tier", boardHandler.CreateTier)

		router.GET("/position/:id", boardHandler.GetPosition)
		router.POST("/position", boardHandler.CreatePosition)
	}
}
