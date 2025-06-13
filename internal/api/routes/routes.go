package routes

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/handlers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, eventService *services.EventsService,
	announcementService *services.AnnouncementService) {
	eventHandler := handlers.NewEventHandler(eventService)
	router.GET("/events", eventHandler.GetEvents)
	router.GET("/events/:id", eventHandler.GetEvent)
	router.POST("/events", eventHandler.CreateEvent)
	router.POST("/events/:id", eventHandler.UpdateEvent)
	router.DELETE("/events/:id", eventHandler.DeleteEvent)
	//announcementService
	announcementHandler := handlers.NewAnnouncementHandler(announcementService)
	router.GET("/announcements/:id", announcementHandler.GetAnnouncement)
	router.POST("/announcements", announcementHandler.CreateAnnouncement)
	router.DELETE("/announcements/:id", announcementHandler.DeleteAnnouncement)
	router.POST("/announcements/:id", announcementHandler.UpdateAnnouncement)
}
