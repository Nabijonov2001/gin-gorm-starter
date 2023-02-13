package services

import (
	"github.com/Nabijonov2001/lfactura/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type client struct {
	db  *gorm.DB
	log *zap.Logger
}

func ClientService(db *gorm.DB, log *zap.Logger) *client {
	return &client{
		db:  db,
		log: log,
	}
}

func (c *client) CreateClient() (models.Client, error) {
	return models.Client{
		Name: "Fazliddin Nabijonov",
		Tin:  "3003030",
	}, nil
}
