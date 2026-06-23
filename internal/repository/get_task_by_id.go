package repository

import (
	"context"
	"final_project/internal/model"
)

func (r *Repository) GetTaskByID(ctx context.Context, id int) (model.Task, error) {
	var task model.Task

	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&task.ID,
		&task.Date,
		&task.Title,
		&task.Comment,
		&task.Repeat,
	)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}
