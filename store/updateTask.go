package store

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (s *Store) UpdateTask(ctx context.Context, task Task) error {
	commandTag, err := s.pool.Exec(ctx,
		`
	UPDATE tasks
	SET title = $1, status = $2
	WHERE ID = $3
	`, task.Title, task.Status, task.ID,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
