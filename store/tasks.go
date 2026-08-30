package store

import (
	"context"
	"fmt"
)

func (s *Store) CreateTask(ctx context.Context, task Task) error {
	_, err := s.pool.Exec(
		ctx,
		`INSERT INTO tasks (id, title, status)
	VALUES ($1, $2, $3)
	`,
		task.ID, task.Title, task.Status,
	)

	return err
}

func (s *Store) GetTasks(ctx context.Context) error {
	commandTag, err := s.pool.Exec(ctx,
		`SELECT * FROM tasks`,
	)

	fmt.Println("commandTag", commandTag)

	return err
}
