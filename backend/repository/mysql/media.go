package mysql

import (
	"context"

	"bacakomik/record/entity"

	"github.com/jackc/pgx/v5"
)

// MediaRepository struct
type MediaRepository struct {
	conn *pgx.Conn
}

// Create Data
// Create Media
func (me *MediaRepository) Create(ctx context.Context, data *entity.Media) (int, error) {
	return 0, nil
}
