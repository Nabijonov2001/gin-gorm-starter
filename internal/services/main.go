package services

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	clientService *client
}

func ServiceHandler(db *gorm.DB, log *zap.Logger) *Service {
	return &Service{
		clientService: ClientService(db, log),
	}
}

func (s *Service) Client() *client {
	return s.clientService
}
