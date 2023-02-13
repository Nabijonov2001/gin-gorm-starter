package main

import (
	"github.com/Nabijonov2001/lfactura/internal/api"
	"github.com/Nabijonov2001/lfactura/internal/db"
	"github.com/Nabijonov2001/lfactura/internal/services"
	"github.com/Nabijonov2001/lfactura/pkg/logger"
	"github.com/spf13/viper"
)

func main() {
	logger := logger.Logger()
	db := db.ConnectDB()

	services := services.ServiceHandler(db, logger)

	server := api.New(&api.RouterOptions{
		Log:     logger,
		Service: *services,
	})

	server.Run(":" + viper.GetString("server.port"))
}
