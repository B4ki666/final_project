package handler

import (
	"net/http"
)

func (h *Handler) GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	response, err := h.Service.GetTaskByID(r.Context(), id)
	if err != nil {
		h.HandleError(err, w)
		return
	}

	WriteJSON(w, http.StatusOK, response, h.Logger)
}
