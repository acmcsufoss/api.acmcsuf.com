package routes

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/handlers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, eventService *services.EventsService, announcementService *services.AnnouncementService) {
	eventsHandler := handlers.NewEventHandler(eventService)
	router.GET("/events", eventsHandler.GetEvents)
	router.GET("/events/:id", eventsHandler.GetEvent)
	router.POST("/events", eventsHandler.CreateEvent)
	router.POST("/events/:id", eventsHandler.UpdateEvent)
	router.DELETE("/events/:id", eventsHandler.DeleteEvent)
	//announcementService
	announcementHandler := handlers.NewAnnouncementHandler(announcementService)
	router.GET("/announcement/:id", announcementHandler.GetAnnouncement)
	router.POST("/announcement", announcementHandler.CreateAnnouncement)
	router.POST("/announcement/:id", announcementHandler.DeleteAnnouncement)
	router.POST("/announcement/:id", announcementHandler.UpdateAnnouncement)
}
