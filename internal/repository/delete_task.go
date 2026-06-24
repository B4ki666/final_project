package repository

import "context"

func (r *Repository) DeleteTask(ctx context.Context, id string) (int, error) {
	query := `DELETE FROM scheduler WHERE id = ?`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rows), nil
}
