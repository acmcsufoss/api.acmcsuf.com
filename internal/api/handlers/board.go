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
	}

	c.JSON(http.StatusOK, officer)
}

func (h *BoardHandler) GetOfficers(c *gin.Context) {
	panic("not implemented")
}

func (h *BoardHandler) UpdateOfficer(c *gin.Context) {
	panic("not implemented")
}

func (h *BoardHandler) DeleteOfficer(c *gin.Context) {
	panic("not implemented")
}

func (h *BoardHandler) CreateOfficer(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.CreateOfficerParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	// TODO: error out if required fields aren't provided
	if err := h.boardService.CreateOfficer(ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create officer",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Officer created successfully",
		"uuid":    params.Uuid,
	})
}

func (h *BoardHandler) GetTier(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to convert string to int",
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
	}

	c.JSON(http.StatusOK, tier)
}

func (h *BoardHandler) CreateTier(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.CreateTierParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	// TODO: error out if required fields aren't provided
	if err := h.boardService.CreateTier(ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create tier",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Tier created successfully",
		"tier":    params.Tier,
	})
}

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
	}

	c.JSON(http.StatusOK, position)
}

func (h *BoardHandler) CreatePosition(c *gin.Context) {
	ctx := c.Request.Context()
	var params models.CreatePositionParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body. " + err.Error(),
		})
		return
	}

	// TODO: error out if required fields aren't provided
	if err := h.boardService.CreatePosition(ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create position",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "Position created successfully",
		"oid":      params.Oid,
		"semester": params.Semester,
		"tier":     params.Tier,
	})
}
