package handler

import (
	"encoding/json"
	"final_project/internal/model"
	"net/http"
)

type GetTasksResponse struct {
	Tasks []model.Task `json:"tasks"`
}

type GetTasksError struct {
	Error string `json:"error"`
}

func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var response GetTasksResponse
	var tasksError GetTasksError
	var tasks []model.Task

	search := r.URL.Query().Get("search")
	if search == "" {
		var err error
		tasks, err = h.Service.GetTasks(r.Context())
		if err != nil {
			tasksError.Error = err.Error()
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(tasksError)
			return
		}
	} else {
		var err error
		tasks, err = h.Service.SearchTasks(r.Context(), search)
		if err != nil {
			tasksError.Error = err.Error()
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(tasksError)
			return
		}
	}

	response.Tasks = tasks

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.Logger.Printf("Error encoding tasks response: %v", err)
	}

}
