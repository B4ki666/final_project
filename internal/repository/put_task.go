package repository

import (
	"context"
	"database/sql"
	"final_project/internal/model"
)

func (r *Repository) PutTask(ctx context.Context, task model.Task) (int, error) {
	query := `UPDATE scheduler SET date = :date, title = :title, comment = :comment, 
	repeat = :repeat WHERE id = :id`

	res, err := r.db.ExecContext(ctx, query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID),
	)

	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rows), nil
}
