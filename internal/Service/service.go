package service

import (
	"context"
	"final_project/internal/model"
	"log"
)

type TODORepository interface {
	AddTask(ctx context.Context, task model.Task) (string, error)
	GetTasks(ctx context.Context) ([]model.Task, error)
	SearchTasksByText(ctx context.Context, search string) ([]model.Task, error)
	SearchTasksByDate(ctx context.Context, date string) ([]model.Task, error)
	GetTaskByID(ctx context.Context, id int) (model.Task, error)
	PutTask(ctx context.Context, task model.Task) (int, error)
}

type Service struct {
	Repo   TODORepository
	Logger *log.Logger
}

func NewService(repo TODORepository, logger *log.Logger) *Service {
	return &Service{
		Repo:   repo,
		Logger: logger,
	}
}
