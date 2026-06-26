package service

import (
	"context"
	"net/http"
	"strconv"
)

func (s *Service) DeleteTask(ctx context.Context, id string) error {
	if id == "" {
		return NewError(http.StatusBadRequest, NoID)
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		return NewError(http.StatusBadRequest, InvalidID)
	}

	rowsAffected, err := s.Repo.DeleteTask(ctx, id)

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return NewError(http.StatusNotFound, NoTask)
	}

	return nil
}
