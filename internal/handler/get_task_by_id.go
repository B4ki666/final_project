package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	response, err := h.Service.GetTaskByID(r.Context(), id)
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
