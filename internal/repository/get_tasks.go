package repository

import (
	"context"
	"database/sql"
	"final_project/internal/model"
)

func (r *Repository) GetTasks(ctx context.Context) ([]model.Task, error) {
	tasks := make([]model.Task, 0)

	query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT 50`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.Date,
			&task.Title,
			&task.Comment,
			&task.Repeat,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repository) SearchTasksByDate(ctx context.Context, date string) ([]model.Task, error) {
	tasks := make([]model.Task, 0)

	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE date = :date LIMIT 50`
	rows, err := r.db.QueryContext(ctx, query, sql.Named("date", date))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.Date,
			&task.Title,
			&task.Comment,
			&task.Repeat,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repository) SearchTasksByText(ctx context.Context, search string) ([]model.Task, error) {
	tasks := make([]model.Task, 0)

	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE title
	 LIKE :search OR comment LIKE :search ORDER BY date LIMIT 50`
	pattern := "%" + search + "%"
	rows, err := r.db.QueryContext(ctx, query, sql.Named("search", pattern))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.Date,
			&task.Title,
			&task.Comment,
			&task.Repeat,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
