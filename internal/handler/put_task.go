package handler

import (
	"encoding/json"
	"final_project/internal/model"
	service "final_project/internal/service"
	"net/http"
)

func (h *Handler) PutTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		h.HandleError(service.NewError(http.StatusBadRequest, DecodeError), w)
		return
	}

	if err := h.Service.PutTask(ctx, task); err != nil {
		h.HandleError(err, w)
		return
	}

	WriteJSON(w, http.StatusOK, struct{}{}, h.Logger)
}
