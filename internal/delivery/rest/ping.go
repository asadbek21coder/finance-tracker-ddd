package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) pong(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Pong",
	})
}
