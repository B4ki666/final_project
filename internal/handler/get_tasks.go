package handler

import (
	"final_project/internal/model"
	"net/http"
)

type GetTasksResponse struct {
	Tasks []model.Task `json:"tasks"`
}

func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []model.Task

	search := r.URL.Query().Get("search")
	if search == "" {
		var err error
		if tasks, err = h.Service.GetTasks(r.Context()); err != nil {
			h.HandleError(err, w)
			return
		}
	} else {
		var err error
		if tasks, err = h.Service.SearchTasks(r.Context(), search); err != nil {
			h.HandleError(err, w)
			return
		}
	}

	WriteJSON(w, http.StatusOK, GetTasksResponse{Tasks: tasks}, h.Logger)
}
