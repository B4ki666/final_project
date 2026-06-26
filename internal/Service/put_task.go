package service

import (
	"context"
	"final_project/internal/model"
	"net/http"
)

func (s *Service) PutTask(ctx context.Context, task model.Task) error {
	if task.Title == "" {
		return NewError(http.StatusBadRequest, NoTitle)
	}

	if task.ID == "" {
		return NewError(http.StatusBadRequest, NoID)
	}

	err := CheckDate(&task)
	if err != nil {
		return NewError(http.StatusBadRequest, err.Error())
	}

	rowsAffected, err := s.Repo.PutTask(ctx, task)

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return NewError(http.StatusNotFound, NoTask)
	}

	return nil
}
