package handler

import (
	"encoding/json"
	service "final_project/internal/Service"
	"final_project/internal/model"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	Logger  *log.Logger
	Service *service.Service
}

type AddTaskResponse struct {
	ID string `json:"id"`
}

type AddTaskError struct {
	Error string `json:"error"`
}

func NewHandler(logger *log.Logger, service *service.Service) *Handler {
	return &Handler{
		Logger:  logger,
		Service: service,
	}
}

func (h *Handler) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	var taskError AddTaskError
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		taskError.Error = fmt.Sprintf("%v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(taskError)
		return
	}

	id, err := h.Service.AddTask(ctx, task)
	if err != nil {
		taskError.Error = fmt.Sprintf("%v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(taskError)
		return
	}

	response := AddTaskResponse{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.Logger.Printf("Error encoding task id: %v", err)
	}
}
