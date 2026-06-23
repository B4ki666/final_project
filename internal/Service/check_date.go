package service

import (
	"final_project/internal/model"
	"final_project/internal/scheduler"
	"time"
)

const (
	format = "20060102"
)

func isPastDate(now, date time.Time) bool {
	nowStr := now.Format(format)
	nowDate, _ := time.Parse(format, nowStr)
	return nowDate.After(date)
}

func CheckDate(task *model.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(format)
	}

	t, err := time.Parse(format, task.Date)
	if err != nil {
		return err
	}

	var next string

	if task.Repeat != "" {
		next, err = scheduler.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if isPastDate(now, t) {
		if task.Repeat == "" {
			task.Date = now.Format(format)
		} else {
			task.Date = next
		}
	}

	return nil
}
