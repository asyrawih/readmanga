package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context, dsn string) (*pgx.Conn, error) {
	c, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	// Check This DB must be Connect to database server
	if err = c.Ping(ctx); err != nil {
		return nil, err
	}

	return c, err
}
