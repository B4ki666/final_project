package repository

import (
	"context"
	"database/sql"
)

func (r *Repository) UpdateTaskDate(ctx context.Context, next string, id string) error {
	query := `UPDATE scheduler SET date = :date WHERE id = :id`

	_, err := r.db.ExecContext(ctx, query,
		sql.Named("date", next),
		sql.Named("id", id),
	)

	if err != nil {
		return err
	}

	return nil
}
