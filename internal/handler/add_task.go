package handler

import (
	"encoding/json"
	"final_project/internal/model"
	service "final_project/internal/service"
	"net/http"
)

type AddTaskResponse struct {
	ID string `json:"id"`
}

func (h *Handler) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		h.HandleError(service.NewError(http.StatusBadRequest, DecodeError), w)
		return
	}

	id, err := h.Service.AddTask(ctx, task)
	if err != nil {
		h.HandleError(err, w)
		return
	}

	WriteJSON(w, http.StatusCreated, AddTaskResponse{ID: id}, h.Logger)
}
