// Package handlers handles http requests and responses
// Business logic belongs in services, not here
package handlers

import (
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/gin-gonic/gin"
)

type AnnouncementHandler struct {
	announcementService *services.AnnouncementService
}

func NewAnnouncementHandler(announcementService *services.AnnouncementService) *AnnouncementHandler {
	return &AnnouncementHandler{announcementService: announcementService}
}

func (h *AnnouncementHandler) GetAnnouncement(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	announcement, err := h.announcementService.GetAnnouncement(ctx, id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Announcement not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Not implemented",
		})
	}

	c.JSON(http.StatusOK, announcement)
}

func (h *AnnouncementHandler) CreateAnnouncement(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.CreateAnnouncementParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	announcement, err := h.announcementService.CreateAnnouncement(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Not implemented",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event created successfully",
		"uuid":    params.Uuid,
	})
	c.JSON(http.StatusOK, announcement)
}
