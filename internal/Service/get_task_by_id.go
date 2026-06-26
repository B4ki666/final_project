package service

import (
	"context"
	"database/sql"
	"errors"
	"final_project/internal/model"
	"net/http"
	"strconv"
)

func (s *Service) GetTaskByID(ctx context.Context, id string) (model.Task, error) {

	if id == "" {
		return model.Task{}, NewError(http.StatusBadRequest, NoID)
	}

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return model.Task{}, NewError(http.StatusBadRequest, InvalidID)
	}

	task, err := s.Repo.GetTaskByID(ctx, taskID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, NewError(http.StatusNotFound, NoTask)
		}

		return model.Task{}, err
	}

	return task, nil
}
