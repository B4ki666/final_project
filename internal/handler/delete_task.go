package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := h.Service.DeleteTask(r.Context(), id); err != nil {
		h.HandleError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(struct{}{}); err != nil {
		h.Logger.Printf("Error encoding tasks response: %v", err)
	}
}
