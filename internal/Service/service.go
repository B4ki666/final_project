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
	UpdateTaskDate(ctx context.Context, next string, id string) error
	DeleteTask(ctx context.Context, id string) (int, error)
}

type Service struct {
	Repo     TODORepository
	Logger   *log.Logger
	Password string
}

func NewService(repo TODORepository, logger *log.Logger, password string) *Service {
	return &Service{
		Repo:     repo,
		Logger:   logger,
		Password: password,
	}
}
