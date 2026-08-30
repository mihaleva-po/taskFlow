package store

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (s *Store) DeleteTask(ctx context.Context, id string) error {
	commandTag, err := s.pool.Exec(ctx,
		`
	DELETE FROM tasks
	WHERE ID = $1
	`,
		id)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
