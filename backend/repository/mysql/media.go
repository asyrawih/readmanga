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

// Delete implements adapter.RepoMediaCreational.
func (*MediaRepository) Delete(ctx context.Context, id int) bool {
	panic("unimplemented")
}

// GetAll implements adapter.RepoMediaCreational.
func (*MediaRepository) GetAll(ctx context.Context) []*entity.Manga {
	panic("unimplemented")
}

// GetOne implements adapter.RepoMediaCreational.
func (*MediaRepository) GetOne(ctx context.Context, id int) *entity.Manga {
	panic("unimplemented")
}

// NewApi implements adapter.RepoMediaCreational.
func (*MediaRepository) NewApi() *MediaRepository {
	panic("unimplemented")
}

// Update implements adapter.RepoMediaCreational.
func (*MediaRepository) Update(ctx context.Context, data *entity.Media, id int) error {
	panic("unimplemented")
}

// NewMediaRepository function
func NewMediaRepository(conn *pgx.Conn) *MediaRepository {
	return &MediaRepository{
		conn: conn,
	}
}

// Create Data
// Create Media
func (me *MediaRepository) Create(ctx context.Context, data *entity.Media) (int, error) {
	var ID int
	sqlString := `INSERT INTO medias (model_type, model_id, url) VALUES ($1, $2, $3) returning id;`
	ct := me.conn.QueryRow(ctx, sqlString, data.ModelType, data.ModelID, data.URL)
	if err := ct.Scan(&ID); err != nil {
		return 0, nil
	}
	return ID, nil
}
