package routes

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/handlers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, eventService *services.EventsService) {
	eventsHandler := handlers.NewEventHandler(eventService)

	router.GET("/events")
	router.GET("/events/:id")
	router.POST("/events")
	router.POST("/events/:id")
	router.DELETE("/events/:id")
	// ss.Get(path, s.Resources())
	// ss.Post(path, s.PostResources())
	// ss.Post(path, s.BatchPostResources())
	// ss.Get(path+"/{id}", s.Resource())
	// ss.Post(path+"/{id}", s.PostResource())
	// ss.Post(path+"/{id}", s.BatchPostResource())
	// ss.Delete(path+"/{id}", s.DeleteResource())

}
