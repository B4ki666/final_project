package handler

import (
	"final_project/internal/scheduler"
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
			h.Logger.Printf("invalid request: %v", err)
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}
	} else {
		now = time.Now()
	}

	result, err := scheduler.NextDate(now, dateStr, repeatStr)
	if err != nil {
		h.Logger.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
