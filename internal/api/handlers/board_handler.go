// Package handlers handles http requests and responses
// Business logic belongs in services, not here
package handlers

import (
	"net/http"

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
