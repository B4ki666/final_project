package handler

import (
	"encoding/json"
	service "final_project/internal/Service"
	"final_project/internal/model"
	"net/http"
)

func (h *Handler) PutTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		h.SendingErrors(service.NewError(400, "invalid json"), w)
		return
	}

	response, err := h.Service.PutTask(ctx, task)
	if err != nil {
		h.SendingErrors(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.Logger.Printf("Error encoding tasks response: %v", err)
	}
}
