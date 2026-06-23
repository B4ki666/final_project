package service

import (
	"context"
	"database/sql"
	"errors"
	"final_project/internal/model"
	"strconv"
)

func (s *Service) GetTaskByID(ctx context.Context, id string) (model.Task, error) {

	if id == "" {
		return model.Task{}, NewError(400, "ID not specified")
	}

	taskid, err := strconv.Atoi(id)
	if err != nil {
		return model.Task{}, NewError(400, "invalid ID")
	}

	task, err := s.Repo.GetTaskByID(ctx, taskid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, NewError(404, "task not found")
		}

		return model.Task{}, err
	}

	return task, nil
}
