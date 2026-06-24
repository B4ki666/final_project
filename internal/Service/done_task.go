package service

import (
	"context"
	"database/sql"
	"errors"
	"final_project/internal/scheduler"
	"strconv"
	"time"
)

func (s *Service) DoneTask(ctx context.Context, id string) error {
	if id == "" {
		return NewError(400, "ID not specified")
	}

	taskid, err := strconv.Atoi(id)
	if err != nil {
		return NewError(400, "invalid ID")
	}

	task, err := s.Repo.GetTaskByID(ctx, taskid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return NewError(404, "task not found")
		}

		return err
	}

	if task.Repeat == "" {
		result, err := s.Repo.DeleteTask(ctx, id)
		if err != nil {
			return err
		}

		if result == 0 {
			return NewError(404, "task not found")
		}

		return nil
	}

	now := time.Now()
	nextDate, err := scheduler.NextDate(now, task.Date, task.Repeat)
	if err != nil {
		return NewError(400, err.Error())
	}

	err = s.Repo.UpdateTaskDate(ctx, nextDate, id)
	if err != nil {
		return err
	}

	return nil
}
