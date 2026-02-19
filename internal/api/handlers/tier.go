// Package handlers handles http requests and responses
// Business logic belongs in services, not here
package handlers

import (
	"net/http"
	"strconv"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	dto_request "github.com/acmcsufoss/api.acmcsuf.com/internal/dto/request"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/mapper"
	"github.com/gin-gonic/gin"
)

type TierHandler struct {
	tierService services.TierServicer
}

func NewTierHandler(tierService services.TierServicer) *TierHandler {
	return &TierHandler{tierService: tierService}
}

// GetTiers godoc
//
//	@Summary		Get all tiers
//	@Description	Get all tiers from the database
//	@Tags			Board
//	@Accept			json
//	@Produce		json
//	@Success		200 {array} dto_request.Tier "List of tiers"
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/tiers [get]
func (h *TierHandler) GetTiers(c *gin.Context) {
	ctx := c.Request.Context()

	tiers, err := h.tierService.List(ctx)
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
//	@Success		200 {object} dbmodels.Tier "Tier details"
//	@Failure		400 {object} map[string]string
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/tiers/{id} [get]
func (h *TierHandler) GetTier(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid tier number",
		})
		return
	}

	tier, err := h.tierService.Get(ctx, int64(id))
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
//	@Param			body body domain.Tier true "Tier data"
//	@Success		200 {object} map[string]interface{} "Success message with tier number"
//	@Failure		400 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/tiers [post]
func (h *TierHandler) CreateTier(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto_request.Tier

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	if err := h.tierService.Create(ctx, mapper.ToTierDomain(&params)); err != nil {
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
//	@Param			body body domain.Tier true "Updated tier data"
//	@Success		200 {object} map[string]string "Success message"
//	@Failure		400 {object} map[string]string
//	@Failure		404 {object} map[string]string
//	@Failure		500 {object} map[string]string
//	@Router			/v1/board/tiers/{id} [put]
func (h *TierHandler) UpdateTier(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto_request.UpdateTier
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

	params.Tier = id

	if err := h.tierService.Update(ctx, int64(id), mapper.ToUpdateTierDomain(&params)); err != nil {
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
func (h *TierHandler) DeleteTier(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid tier number",
		})
		return
	}

	if err := h.tierService.Delete(ctx, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete tier",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tier deleted successfully",
	})
}
