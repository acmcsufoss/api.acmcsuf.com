package handlers

import (
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/sqlite"
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
	// events := h.eventsService.GetEvent()
	// c.JSON(http.StatusOK, events)
}
