// Package handlers handles http requests and responses
// Business logic belongs in services, not here
package handlers

import (
	"net/http"
	"strconv"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/gin-gonic/gin"
)

type BoardHandler struct {
	boardService services.BoardServicer
}

func NewBoardHandler(boardService services.BoardServicer) *BoardHandler {
	return &BoardHandler{boardService: boardService}
}

// GetOfficer godoc
//
//	@Summary		Get an Officer by UUID
//	@Description	Retrieves a single officer from the database.
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			id path string true "Officer UUID"
//	@Success		200 {object} models.Officer "Officer details"
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/officers/{id} [get]
func (h *BoardHandler) GetOfficer(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	officer, err := h.boardService.GetOfficer(ctx, id)

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
//	@Success		200 {array} models.Officer "List of officers"
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/officers [get]
func (h *BoardHandler) GetOfficers(c *gin.Context) {
	ctx := c.Request.Context()

	officers, err := h.boardService.ListOfficers(ctx)
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
//	@Param			body body models.CreateOfficerParams true "Officer data"
//	@Success		200 {object} map[string]interface{} "Success message with UUID"
//	@Failure		400 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/officers [post]
func (h *BoardHandler) CreateOfficer(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.CreateOfficerParams

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

	if err := h.boardService.CreateOfficer(ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create officer. " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Officer created successfully",
		"uuid":    params.Uuid,
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
//	@Param			body body models.UpdateOfficerParams true "Updated officer data"
//	@Success		200 {object} map[string]string "Success message"
//	@Failure		400 {object} map[string]string
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/officers/{id} [put]
func (h *BoardHandler) UpdateOfficer(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.UpdateOfficerParams
	id := c.Param("id")

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if err := h.boardService.UpdateOfficer(ctx, id, params); err != nil {
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
func (h *BoardHandler) DeleteOfficer(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	if err := h.boardService.DeleteOfficer(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete officer",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Officer deleted successfully",
	})
}

// GetTiers godoc
//
//	@Summary		Get all tiers
//	@Description	Get all tiers from the database
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Success		200 {array} models.Tier "List of tiers"
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/tiers [get]
func (h *BoardHandler) GetTiers(c *gin.Context) {
	ctx := c.Request.Context()

	tiers, err := h.boardService.ListTiers(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve tiers",
		})
		return
	}

	c.JSON(http.StatusOK, tiers)
}

// GetTier godoc
//
//	@Summary		Get a Tier by tier number
//	@Description	Retrieves a single tier from the database.
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			id path int true "Tier number"
//	@Success		200 {object} models.Tier "Tier details"
//	@Failure		400 {object} map[string]string
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/tiers/{id} [get]
func (h *BoardHandler) GetTier(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid tier number",
		})
		return
	}

	tier, err := h.boardService.GetTier(ctx, int64(id))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Tier not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve tier",
		})
		return
	}

	c.JSON(http.StatusOK, tier)
}

// CreateTier godoc
//
//	@Summary		Creates a new tier
//	@Description	Creates a new tier in the database.
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			body body models.CreateTierParams true "Tier data"
//	@Success		200 {object} map[string]interface{} "Success message with tier number"
//	@Failure		400 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/tiers [post]
func (h *BoardHandler) CreateTier(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.CreateTierParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if err := h.boardService.CreateTier(ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create tier. " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tier created successfully",
		"tier":    params.Tier,
	})
}

// UpdateTier godoc
//
//	@Summary		Updates a tier
//	@Description	Updates a tier in the database
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			id path int true "Tier number"
//	@Param			body body models.UpdateTierParams true "Updated tier data"
//	@Success		200 {object} map[string]string "Success message"
//	@Failure		400 {object} map[string]string
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/tiers/{id} [put]
func (h *BoardHandler) UpdateTier(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.UpdateTierParams
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid tier number",
		})
		return
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	params.Tier = int64(id)

	if err := h.boardService.UpdateTier(ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update tier. " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tier updated successfully",
		"tier":    params.Tier,
	})
}

// DeleteTier godoc
//
//	@Summary		Deletes a tier
//	@Description	Delete a tier from the database
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Param			id path int true "Tier number"
//	@Success		200 {object} map[string]string "Success message"
//	@Failure		400 {object} map[string]string
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/tiers/{id} [delete]
func (h *BoardHandler) DeleteTier(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid tier number",
		})
		return
	}

	if err := h.boardService.DeleteTier(ctx, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete tier",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tier deleted successfully",
	})
}

// GetPositions godoc
//
//	@Summary		Get all positions
//	@Description	Get all positions from the database
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Success		200 {array} models.Position "List of positions"
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/positions [get]
func (h *BoardHandler) GetPositions(c *gin.Context) {
	ctx := c.Request.Context()

	positions, err := h.boardService.ListPositions(ctx)
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
//	@Success		200 {object} models.Position "Position details"
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/positions/{id} [get]
func (h *BoardHandler) GetPosition(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	position, err := h.boardService.GetPosition(ctx, id)

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
//	@Param			body body models.CreatePositionParams true "Position data"
//	@Success		200 {object} map[string]interface{} "Success message"
//	@Failure		400 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/positions [post]
func (h *BoardHandler) CreatePosition(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.CreatePositionParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if err := h.boardService.CreatePosition(ctx, params); err != nil {
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
//	@Param			body body models.UpdatePositionParams true "Updated position data (must include oid, semester, tier)"
//	@Success		200 {object} map[string]string "Success message"
//	@Failure		400 {object} map[string]string
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/positions [put]
func (h *BoardHandler) UpdatePosition(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.UpdatePositionParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if err := h.boardService.UpdatePosition(ctx, params); err != nil {
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
//	@Param			body body models.DeletePositionParams true "Position identifier"
//	@Success		200 {object} map[string]string "Success message"
//	@Failure		400 {object} map[string]string
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/positions [delete]
func (h *BoardHandler) DeletePosition(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.DeletePositionParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if err := h.boardService.DeletePosition(ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete position",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Position deleted successfully",
	})
}
