// This file (v1.go) initializes v1 routes used by the server. Called by server.go

package routes

import (
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/handlers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/gin-gonic/gin"
)

func SetupV1Routes(router *gin.Engine, eventService services.EventsServicer,
	announcementService services.AnnouncementServicer) {
	router.GET("/swagger/*any", handlers.NewSwaggerHandler())
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Version 1 routes
	v1 := router.Group("/v1")
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

	}
}
