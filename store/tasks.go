package store

import (
	"context"

	"github.com/jackc/pgx/v5"
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

func (s *Store) GetTasks(ctx context.Context) ([]Task, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT * FROM tasks`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&task.ID, &task.Title, &task.Status,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *Store) GetTask(ctx context.Context, id string) (Task, error) {

	var task Task

	err := s.pool.QueryRow(ctx,
		`SELECT * FROM tasks
WHERE ID = $1
`, id,
	).Scan(&task.ID, &task.Status, &task.Title)

	if err != nil {
		return Task{}, err
	}

	return task, nil

}

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
