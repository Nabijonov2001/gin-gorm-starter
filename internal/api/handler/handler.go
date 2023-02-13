package handler

import (
	"github.com/Nabijonov2001/lfactura/internal/services"
	"go.uber.org/zap"
)

type Handler struct {
	log     *zap.Logger
	service services.Service
}

type HandlerOptions struct {
	Log     *zap.Logger
	Service services.Service
}

func NewHandler(opts *HandlerOptions) *Handler {
	return &Handler{
		log:     opts.Log,
		service: opts.Service,
	}
}
