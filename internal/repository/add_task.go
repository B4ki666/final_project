package repository

import (
	"context"
	"database/sql"
	"final_project/internal/model"
	"fmt"
)

func (r *Repository) AddTask(ctx context.Context, task model.Task) (string, error) {
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)`
	res, err := r.db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)
	if err != nil {
		return "", err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", id), nil
}
