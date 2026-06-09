package handler

import "log"

type Handler struct {
	Logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		Logger: logger,
	}
}
