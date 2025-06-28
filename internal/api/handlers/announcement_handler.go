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
	announcementService services.AnnouncementServicer
}

func NewAnnouncementHandler(announcementService services.AnnouncementServicer) *AnnouncementHandler {
	return &AnnouncementHandler{announcementService: announcementService}
}

// GetAnnouncement godoc
//
//	@Summary		Get an announcement by ID
//	@Description	Retrieves a single announcement from the database.
//	@Tags			Announcements
//	@Accept			json
//	@Produce		json
//	@Param			id path string true "Announcement ID"
//	@Success		200 {object} models.Announcement "Announcement details"
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/announcements/{id} [get]
func (h *AnnouncementHandler) GetAnnouncement(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	announcement, err := h.announcementService.Get(ctx, id)

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

// CreateAnnouncement godoc
//
//	@Summary		Create new Announcement
//	@Description	Creates a new announcement and generates unique ID
//	@Tags			Announcements
//	@Accept			json
//	@Produce		json
//	@Param			body body models.CreateAnnouncementParams true "Announcement data"
//	@Success		200 {object} map[string]interface{} "Success message with UUID"
//	@Failure		400 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/announcements [post]
func (h *AnnouncementHandler) CreateAnnouncement(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.CreateAnnouncementParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if err := h.announcementService.Create(ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Not implemented",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Announcement created successfully",
		"uuid":    params.Uuid,
	})
}

// UpdateAnnouncement godoc
//
// @Summary		Updates the Announcement of Choice
// @Description	Updates the Announcement of choice in the database
// @Tags		Announcements
// @Accept		json
// @Produce		json
// @Param		id path string true "Announcement ID"
// @Param		body body models.UpdateAnnouncementParams true "Updated announcement data"
// @Success		200 {object} map[string]string "Success message"
// @Failure		400 {object} map[string]string
// @Failure		404 {object} map[string]string
// @Failure		500 {object} map[string]string
// @Router		/announcements/{id} [put]
func (h *AnnouncementHandler) UpdateAnnouncement(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.UpdateAnnouncementParams
	id := c.Param("id")

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
	}

	if err := h.announcementService.Update(ctx, id, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update announcement",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Announcement updated successfully",
		"uuid":    params.Uuid,
	})
}

// DeleteAnnouncement godoc
//
//		@Summary		Deletes the Announcement of Choice
//		@Description	Deletes the Announcement of choice in the database
//		@Tags			Announcements
//		@Accept			json
//		@Produce		json
//	 	@Param			id path string true "Announcement ID"
//		@Success		200 {object} map[string]string "Success message"
//		@Failure		404 {object} map[string]string
//		@Failure		500 {object} map[string]string
//		@Router			/announcements/{id} [delete]
func (h *AnnouncementHandler) DeleteAnnouncement(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	if err := h.announcementService.Delete(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete announcement",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Announcement updated successfully",
	})
}
