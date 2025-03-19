package routes

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/handlers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, eventService *services.EventsService) {
	eventsHandler := handlers.NewEventHandler(eventService)

	// TODO: implement functions in events_handler.go and uncomment these

	// router.GET("/events", eventsHandler.GetEvents)
	router.GET("/events/:id", eventsHandler.GetEvent)
	// router.POST("/events", eventsHandler.CreateEvent)
	// router.POST("/events/:id", eventsHandler.UpdateEvent)
	// router.DELETE("/events/:id", eventsHandler.DeleteEvent)
}
