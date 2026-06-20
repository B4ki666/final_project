package handler

import (
	"encoding/json"
	"final_project/internal/model"
	"net/http"
)

type AddTaskResponse struct {
	ID string `json:"id"`
}

type AddTaskError struct {
	Error string `json:"error"`
}

func (h *Handler) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	var taskError AddTaskError
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		taskError.Error = err.Error()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(taskError)
		return
	}

	id, err := h.Service.AddTask(ctx, task)
	if err != nil {
		taskError.Error = err.Error()
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
