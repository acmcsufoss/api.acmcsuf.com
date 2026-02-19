// Package handlers handles http requests and responses
// Business logic belongs in services, not here
package handlers

import (
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	dto_request "github.com/acmcsufoss/api.acmcsuf.com/internal/dto/request"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/mapper"
	"github.com/gin-gonic/gin"
)

type OfficerHandler struct {
	officerService services.OfficerServicer
}

func NewOfficerHandler(officerService services.OfficerServicer) *OfficerHandler {
	return &OfficerHandler{officerService: officerService}
}

// GetOfficer godoc
//
//	@Summary		Get an Officer by UUID
//	@Description	Retrieves a single officer from the database.
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			id path string true "Officer UUID"
//	@Success		200 {object} dto_request.Officer "Officer details"
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/officers/{id} [get]
func (h *OfficerHandler) GetOfficer(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	officer, err := h.officerService.Get(ctx, id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Officer not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve officer",
		})
		return
	}

	c.JSON(http.StatusOK, officer)
}

// GetOfficers godoc
//
//	@Summary		Get all officers
//	@Description	Get all officers from the database
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Success		200 {array} domain.Officer "List of officers"
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/officers [get]
func (h *OfficerHandler) GetOfficers(c *gin.Context) {
	ctx := c.Request.Context()

	officers, err := h.officerService.List(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve officers",
		})
		return
	}

	c.JSON(http.StatusOK, officers)
}

// CreateOfficer godoc
//
//	@Summary		Creates a new officer
//	@Description	Creates a new officer in the database.
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			body body domain.Officer true "Officer data"
//	@Success		200 {object} map[string]interface{} "Success message with UUID"
//	@Failure		400 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/officers [post]
func (h *OfficerHandler) CreateOfficer(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto_request.Officer

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if params.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FullName is a required field",
		})
		return
	}

	if err := h.officerService.Create(ctx, mapper.ToOfficerDomain(&params)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create officer. " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Officer created successfully",
		"uuid":    mapper.ToOfficerDomain(&params).Uuid,
	})
}

// UpdateOfficer godoc
//
//	@Summary		Updates an officer
//	@Description	Updates an officer in the database
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			id path string true "Officer UUID"
//	@Param			body body domain.Officer true "Updated officer data"
//	@Success		200 {object} map[string]string "Success message"
//	@Failure		400 {object} map[string]string
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/officers/{id} [put]
func (h *OfficerHandler) UpdateOfficer(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto_request.UpdateOfficer
	id := c.Param("id")

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if err := h.officerService.Update(ctx, id, mapper.ToUpdateOfficerDomain(&params)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update officer. " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Officer updated successfully",
		"uuid":    id,
	})
}

// DeleteOfficer godoc
//
//	@Summary		Deletes an officer
//	@Description	Delete an officer from the database
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			id path string true "Officer UUID"
//	@Success		200 {object} map[string]string "Success message"
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/officers/{id} [delete]
func (h *OfficerHandler) DeleteOfficer(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	if err := h.officerService.Delete(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete officer",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Officer deleted successfully",
	})
}
