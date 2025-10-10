// This file (v1.go) initializes v1 routes used by the server. Called by server.go

package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/handlers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
)

func SetupV1(router *gin.Engine, eventService services.EventsServicer,
	announcementService services.AnnouncementServicer) {

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
