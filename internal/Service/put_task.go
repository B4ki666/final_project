package service

import (
	"context"
	"final_project/internal/model"
)

func (s *Service) PutTask(ctx context.Context, task model.Task) (model.Task, error) {
	if task.Title == "" {
		return model.Task{}, NewError(400, "task title not specified")
	}

	if task.ID == "" {
		return model.Task{}, NewError(400, "ID not specified")
	}

	err := CheckDate(&task)
	if err != nil {
		return model.Task{}, NewError(400, err.Error())
	}

	result, err := s.Repo.PutTask(ctx, task)

	if err != nil {
		return model.Task{}, err
	}

	if result == 0 {
		return model.Task{}, NewError(404, "task not found")
	}

	return model.Task{}, nil
}
