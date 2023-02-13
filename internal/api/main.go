package api

import (
	"github.com/Nabijonov2001/lfactura/internal/api/handler"
	"github.com/Nabijonov2001/lfactura/internal/api/routes"
	"github.com/Nabijonov2001/lfactura/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RouterOptions struct {
	Log     *zap.Logger
	Service services.Service
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())

	handler := handler.NewHandler(&handler.HandlerOptions{
		Log:     opt.Log,
		Service: opt.Service,
	})

	api := router.Group("/v1")

	// all api endpoints
	routes.ClientRoutes(api, handler)

	return router

}
