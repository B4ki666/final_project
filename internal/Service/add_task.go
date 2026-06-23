package service

import (
	"context"
	"errors"
	"final_project/internal/model"
)

func (s *Service) AddTask(ctx context.Context, task model.Task) (string, error) {
	if task.Title == "" {
		return "", errors.New("task title not specified")
	}

	err := CheckDate(&task)
	if err != nil {
		return "", err
	}

	return s.Repo.AddTask(ctx, task)
}
