package routes

import (
	"github.com/Nabijonov2001/lfactura/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func ClientRoutes(r *gin.RouterGroup, h *handler.Handler) {
	r.POST("/client/create", h.CreateClient)
}
