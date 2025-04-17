// Package handlers handles http requests and responses
// Business logic belongs in services, not here
package handlers

import (
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/gin-gonic/gin"
)

type EventsHandler struct {
	eventsService *services.EventsService
}

func NewEventHandler(eventService *services.EventsService) *EventsHandler {
	return &EventsHandler{
		eventsService: eventService,
	}
}

func (h *EventsHandler) GetEvent(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	event, err := h.eventsService.Get(ctx, id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Event not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve event",
		})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (h *EventsHandler) CreateEvent(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.CreateEventParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if params.Location == "" || params.Host == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Location and Host are required fields",
		})
		return
	}

	err := h.eventsService.Create(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create event. " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event created successfully",
		"uuid":    params.Uuid,
	})
}

func (h *EventsHandler) GetEvents(c *gin.Context) {
	ctx := c.Request.Context()
	host := c.Query("host")
	filters := []any{}

	if host != "" {
		filters = append(filters, services.HostFilter{Host: host})
	}

	events, err := h.eventsService.List(ctx, filters...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve events",
		})
		return
	}
	c.JSON(http.StatusOK, events)
}

func (h *EventsHandler) UpdateEvent(c *gin.Context) {
	// ctx := c.Request.Context()
	// var params models.UpdateEventParams
	panic("implement me")
}

func (h *EventsHandler) DeleteEvent(c *gin.Context) {
	panic("implement me")
}
