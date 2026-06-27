package service

import (
	"context"
	"final_project/internal/model"
	"net/http"
)

func (s *Service) AddTask(ctx context.Context, task model.Task) (string, error) {
	if task.Title == "" {
		return "", NewError(http.StatusBadRequest, NoTitle)
	}

	if err := CheckDate(&task); err != nil {
		return "", err
	}

	return s.Repo.AddTask(ctx, task)
}
