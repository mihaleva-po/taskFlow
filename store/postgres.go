package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	pool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		pool: pool,
	}
}

func NewPostgresPool(ctx context.Context, databaseUrl string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, databaseUrl)

	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
