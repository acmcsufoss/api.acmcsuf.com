// Package handlers handles http requests and responses
// Business logic belongs in services, not here
package handlers

import (
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	dto_request "github.com/acmcsufoss/api.acmcsuf.com/internal/dto/request"
	"github.com/gin-gonic/gin"
)

type PositionHandler struct {
	positionService services.PositionServicer
}

func NewPositionHandler(positionService services.PositionServicer) *PositionHandler {
	return &PositionHandler{positionService: positionService}
}

// GetPositions godoc
//
//	@Summary		Get all positions
//	@Description	Get all positions from the database
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Success		200 {array} dbmodels.Position "List of positions"
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/positions [get]
func (h *PositionHandler) GetPositions(c *gin.Context) {
	ctx := c.Request.Context()

	positions, err := h.positionService.List(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve positions",
		})
		return
	}

	c.JSON(http.StatusOK, positions)
}

// GetPosition godoc
//
//	@Summary		Get a Position by UUID
//	@Description	Retrieves a single position from the database by officer UUID.
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			id path string true "Officer full name"
//	@Success		200 {object} dbmodels.Position "Position details"
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/positions/{id} [get]
func (h *PositionHandler) GetPosition(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	position, err := h.positionService.Get(ctx, id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Position not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve position",
		})
		return
	}

	c.JSON(http.StatusOK, position)
}

// CreatePosition godoc
//
//	@Summary		Creates a new position
//	@Description	Creates a new position in the database.
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			body body dbmodels.CreatePositionParams true "Position data"
//	@Success		200 {object} map[string]interface{} "Success message"
//	@Failure		400 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/positions [post]
func (h *PositionHandler) CreatePosition(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto_request.Position

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if err := h.positionService.Create(ctx, params.ToDomain()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create position. " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Position created successfully",
		"oid":      params.Oid,
		"semester": params.Semester,
		"tier":     params.Tier,
	})
}

// UpdatePosition godoc
//
//	@Summary		Updates a position
//	@Description	Updates a position in the database
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			body body dbmodels.UpdatePositionParams true "Updated position data (must include oid, semester, tier)"
//	@Success		200 {object} map[string]string "Success message"
//	@Failure		400 {object} map[string]string
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/positions [put]
func (h *PositionHandler) UpdatePosition(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto_request.Position
	id := c.Param("id")

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if err := h.positionService.Update(ctx, id, params.ToDomain()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update position. " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Position updated successfully",
		"oid":      params.Oid,
		"semester": params.Semester,
		"tier":     params.Tier,
	})
}

// DeletePosition godoc
//
//	@Summary		Deletes a position
//	@Description	Delete a position from the database (requires oid, semester, and tier)
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			body body dbmodels.DeletePositionParams true "Position identifier"
//	@Success		200 {object} map[string]string "Success message"
//	@Failure		400 {object} map[string]string
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/positions [delete]
func (h *PositionHandler) DeletePosition(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto_request.Position

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if err := h.positionService.DeletePosition(ctx, params.ToDomain()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete position",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Position deleted successfully",
	})
}
