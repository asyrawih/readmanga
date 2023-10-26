package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	cc, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	cc.MaxConns = 100
	cc.MaxConnIdleTime = 5 * time.Minute

	p, err := pgxpool.NewWithConfig(ctx, cc)
	if err != nil {
		return nil, err
	}

	return p, err
}
