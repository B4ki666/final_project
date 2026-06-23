package handler

import (
	"encoding/json"
	"errors"
	service "final_project/internal/Service"
	"net/http"
)

type TaskError struct {
	Error string `json:"error"`
}

func (h *Handler) SendingErrors(err error, w http.ResponseWriter) {
	var taskError TaskError
	var httpErr *service.AppError
	if errors.As(err, &httpErr) {
		taskError.Error = httpErr.Message
		switch httpErr.StatusCode {
		case 400:
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(taskError)
		case 404:
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(taskError)

		default:
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(taskError)
		}
		return
	}

	taskError.Error = err.Error()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(taskError)
}
