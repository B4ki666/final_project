package service

import (
	"context"
	"database/sql"
	"errors"
	"final_project/internal/scheduler"
	"net/http"
	"strconv"
	"time"
)

func (s *Service) DoneTask(ctx context.Context, id string) error {
	if id == "" {
		return NewError(http.StatusBadRequest, NoID)
	}

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return NewError(http.StatusBadRequest, InvalidID)
	}

	task, err := s.Repo.GetTaskByID(ctx, taskID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return NewError(http.StatusNotFound, NoTask)
		}

		return err
	}

	if task.Repeat == "" {
		rowsAffected, err := s.Repo.DeleteTask(ctx, id)
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return NewError(http.StatusNotFound, NoTask)
		}

		return nil
	}

	now := time.Now()
	nextDate, err := scheduler.NextDate(now, task.Date, task.Repeat)
	if err != nil {
		return NewError(http.StatusBadRequest, err.Error())
	}

	if err := s.Repo.UpdateTaskDate(ctx, nextDate, id); err != nil {
		return err
	}

	return nil
}
