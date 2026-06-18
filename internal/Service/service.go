package service

import (
	"context"
	"errors"
	"final_project/internal/model"
	"final_project/internal/scheduler"

	"log"
	"time"
)

const (
	format = "20060102"
)

type TODORepository interface {
	AddTask(ctx context.Context, task model.Task) (string, error)
}

type Service struct {
	Repo   TODORepository
	Logger *log.Logger
}

func NewService(repo TODORepository, logger *log.Logger) *Service {
	return &Service{
		Repo:   repo,
		Logger: logger,
	}
}

func isPastDate(now, date time.Time) bool {
	nowStr := now.Format(format)
	nowDate, _ := time.Parse(format, nowStr)
	return nowDate.After(date)
}

func checkDate(task *model.Task) error {
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

func (s *Service) AddTask(ctx context.Context, task model.Task) (string, error) {
	if task.Title == "" {
		return "", errors.New("task title not specified")
	}

	err := checkDate(&task)
	if err != nil {
		return "", err
	}

	return s.Repo.AddTask(ctx, task)
}
