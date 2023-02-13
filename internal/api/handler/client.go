package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateClient(c *gin.Context) {
	res, _ := h.service.Client().CreateClient()

	c.JSON(200, gin.H{
		"data": res,
	})
}
