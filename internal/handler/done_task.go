package handler

import (
	"net/http"
)

func (h *Handler) DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := h.Service.DoneTask(r.Context(), id); err != nil {
		h.HandleError(err, w)
		return
	}

	WriteJSON(w, http.StatusOK, struct{}{}, h.Logger)
}
