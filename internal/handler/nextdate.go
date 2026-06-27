package handler

import (
	"final_project/internal/scheduler"
	service "final_project/internal/service"
	"net/http"
	"time"
)

const (
	format = "20060102"
)

func (h *Handler) NextDateHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	nowStr := queryParams.Get("now")
	dateStr := queryParams.Get("date")
	repeatStr := queryParams.Get("repeat")

	var now time.Time

	if nowStr != "" {
		var err error
		now, err = time.Parse(format, nowStr)
		if err != nil {
			h.HandleError(service.NewError(http.StatusBadRequest, InvalidDate), w)
			return
		}
	} else {
		now = time.Now()
	}

	response, err := scheduler.NextDate(now, dateStr, repeatStr)
	if err != nil {
		h.HandleError(service.NewError(http.StatusBadRequest, err.Error()), w)
		return
	}

	WriteJSON(w, http.StatusOK, response, h.Logger)
}
