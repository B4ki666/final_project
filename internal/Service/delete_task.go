package service

import (
	"context"
	"strconv"
)

func (s *Service) DeleteTask(ctx context.Context, id string) error {
	if id == "" {
		return NewError(400, "ID not specified")
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		return NewError(400, "invalid ID")
	}

	rowsAffected, err := s.Repo.DeleteTask(ctx, id)

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return NewError(404, "task not found")
	}

	return nil
}
