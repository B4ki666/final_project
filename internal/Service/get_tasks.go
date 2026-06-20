package service

import (
	"context"
	"final_project/internal/model"
	"time"
)

const (
	searchFormat = "02.01.2006"
)

func (s *Service) GetTasks(ctx context.Context) ([]model.Task, error) {
	return s.Repo.GetTasks(ctx)
}

func (s *Service) SearchTasks(ctx context.Context, search string) ([]model.Task, error) {
	t, err := time.Parse(searchFormat, search)

	if err == nil {
		return s.Repo.SearchTasksByDate(ctx, t.Format(format))
	}
	return s.Repo.SearchTasksByText(ctx, search)

}
