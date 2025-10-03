package routes

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/handlers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, eventService services.EventsServicer,
	announcementService services.AnnouncementServicer, boardService services.BoardServicer) {

	eventHandler := handlers.NewEventHandler(eventService)
	router.GET("/events", eventHandler.GetEvents)
	router.GET("/events/:id", eventHandler.GetEvent)
	router.POST("/events", eventHandler.CreateEvent)
	router.PUT("/events/:id", eventHandler.UpdateEvent)
	router.DELETE("/events/:id", eventHandler.DeleteEvent)

	announcementHandler := handlers.NewAnnouncementHandler(announcementService)
	router.GET("/announcements", announcementHandler.GetAnnouncements)
	router.GET("/announcements/:id", announcementHandler.GetAnnouncement)
	router.POST("/announcements", announcementHandler.CreateAnnouncement)
	router.PUT("/announcements/:id", announcementHandler.UpdateAnnouncement)
	router.DELETE("/announcements/:id", announcementHandler.DeleteAnnouncement)

	boardHandler := handlers.NewBoardHandler(boardService)
	// router.GET("/tier/:id", boardHandler.GetTier)
	// router.GET("/position/:id", boardHandler.GetPosition)
	router.GET("/officer", boardHandler.GetOfficers)
	router.GET("/officer/:id", boardHandler.GetOfficer)
	router.POST("/officer/", boardHandler.CreateOfficer)
	router.PUT("/officer/:id", boardHandler.UpdateOfficer)
	router.DELETE("/officer/:id", boardHandler.DeleteOfficer)
}
