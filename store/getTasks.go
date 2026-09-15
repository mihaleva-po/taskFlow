package store

import "context"

func (s *Store) GetTasks(ctx context.Context) ([]Task, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT id, title, status, created_at FROM tasks`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&task.ID, &task.Title, &task.Status, &task.CreatedAT,
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
		`
		SELECT id, title, status, created_at FROM tasks
		WHERE ID = $1
`, id,
	).Scan(&task.ID, &task.Title, &task.Status, &task.CreatedAT)

	if err != nil {
		return Task{}, err
	}

	return task, nil

}
