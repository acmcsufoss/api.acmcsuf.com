// Package handlers handles http requests and responses
// Business logic belongs in services, not here
package handlers

import (
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/dto"
	"github.com/gin-gonic/gin"
)

type EventsHandler struct {
	eventsService services.EventsServicer
}

func NewEventHandler(eventService services.EventsServicer) *EventsHandler {
	return &EventsHandler{
		eventsService: eventService,
	}
}

// GetEvent godoc
//
//	@Summary		Get an Event by ID
//	@Description	Retrieves a single event from the database.
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Param			id path string true "Event ID"
//	@Success		200 {object} dto.Event "Event details"
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/events/{id} [get]
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

	c.JSON(http.StatusOK, dto.EventDomainToDto(&event))
}

// CreateEvent godoc
//
//	@Summary		Creates a new event and generates new ID
//	@Description	Creates a new event in the database.
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Param			body body dto.Event true "Event data"
//	@Success		200 {object} map[string]interface{} "Success message with UUID"
//	@Failure		400 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/events [post]
func (h *EventsHandler) CreateEvent(c *gin.Context) {
	ctx := c.Request.Context()
	var body dto.Event

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	err := h.eventsService.Create(ctx, body.ToDomain())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create event. " + err.Error(),
		})
		return
	}
	// TODO: Get event and return it here. Issue is that the generated sqlc func
	// isn't returning the resource, so the service isn't either.
	c.JSON(http.StatusOK, gin.H{
		"message": "Event created successfully",
	})
}

// GetEvents godoc
//
//	@Summary		Get all the events
//	@Description	Get all the events from the event database
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Param			host query string false "Filter by host"
//	@Success		200 {array} dto.Event "List of events"
//	@Failure		500 {object} map[string]string
//	@Router			/v1/events [get]
func (h *EventsHandler) GetEvents(c *gin.Context) {
	ctx := c.Request.Context()
	host := c.Query("host")
	filters := []any{}

	if host != "" {
		filters = append(filters, &services.HostFilter{Host: host})
	}

	events, err := h.eventsService.List(ctx, filters...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve events",
		})
		return
	}

	dtoEvents := make([]dto.Event, len(events))
	for i, event := range events {
		dtoEvents[i] = dto.EventDomainToDto(&event)
	}
	c.JSON(http.StatusOK, dtoEvents)
}

// UpdateEvent godoc
//
//		@Summary		Updates the Event of Choice
//		@Description	Updates the event of choice in the database
//		@Tags			Events
//		@Accept			json
//		@Produce		json
//	 	@Param			id path string true "Event ID"
//	 	@Param			body body dto.UpdateEvent true "Updated event data"
//		@Success		200 {object} map[string]string "Success message"
//		@Failure		400 {object} map[string]string
//		@Failure		404 {object} map[string]string
//		@Failure		500 {object} map[string]string
//		@Router			/v1/events/{id} [put]
func (h *EventsHandler) UpdateEvent(c *gin.Context) {
	ctx := c.Request.Context()
	var body dto.UpdateEvent
	id := c.Param("id")
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if err := h.eventsService.Update(ctx, id, body.ToDomain()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}

// DeleteEvent godoc
//
//		@Summary		Deletes the Event of Choice
//		@Description	Delete the event of choice from the database
//		@Tags			Events
//		@Accept			json
//		@Produce		json
//	 	@Param			id path string true "Event ID"
//		@Success		200 {object} map[string]string "Success message"
//		@Failure		404 {object} map[string]string
//		@Failure		500 {object} map[string]string
//		@Router			/v1/events/{id} [delete]
func (h *EventsHandler) DeleteEvent(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if err := h.eventsService.Delete(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete event",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
