package handler

import (
	service "final_project/internal/Service"
	"log"
)

type Handler struct {
	Logger  *log.Logger
	Service *service.Service
}

func NewHandler(logger *log.Logger, service *service.Service) *Handler {
	return &Handler{
		Logger:  logger,
		Service: service,
	}
}
