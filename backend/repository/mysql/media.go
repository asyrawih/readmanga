package mysql

import (
	"context"
	"errors"

	"bacakomik/record/entity"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

// MediaRepository struct
type MediaRepository struct {
	conn *pgxpool.Pool
}

// Delete implements adapter.RepoMediaCreational.
func (md *MediaRepository) Delete(ctx context.Context, id int) bool {
	var ID int
	sqlString := `delete from medias as md where md.id = $1;`
	ct := md.conn.QueryRow(ctx, sqlString, id)
	if err := ct.Scan(&ID); err != nil {
		return false
	}
	return true
}

// GetAll implements adapter.RepoMediaCreational.
func (md *MediaRepository) GetAll(ctx context.Context) []*entity.Media {
	var medias []*entity.Media
	sqlString := `select id , model_type , model_id , url from medias limit 200`
	if err := pgxscan.Select(ctx, md.conn, &medias, sqlString); err != nil {
		log.Err(err).Msg("")
		return nil
	}
	return medias
}

// GetOne implements adapter.RepoMediaCreational.
func (md *MediaRepository) GetOne(ctx context.Context, id int) *entity.Media {
	m := new(entity.Media)
	sqlString := `select id , model_type , model_id , url from medias as m where m.id = $1 `
	r := md.conn.QueryRow(ctx, sqlString, id)
	if err := r.Scan(&m.ID, &m.ModelType, &m.ModelID, &m.URL); err != nil {
		log.Err(err).Msg("[mysql](GetOne): fail get scann data")
		return nil
	}
	return m
}

// NewApi implements adapter.RepoMediaCreational.
func (md *MediaRepository) NewApi() *MediaRepository {
	return md
}

// Update implements adapter.RepoMediaCreational.
func (md *MediaRepository) Update(ctx context.Context, data *entity.Media, id int) error {
	sqlString := ` UPDATE medias SET model_type = $1, model_id = $2, url = $3 WHERE medias.id = $4;`
	ct, err := md.conn.Exec(ctx, sqlString, data.ModelType, data.ModelID, data.URL, id)
	if err != nil {
		return err
	}
	if b := ct.Update(); !b {
		return errors.New("fail update")
	}

	return nil
}

// NewMediaRepository function
func NewMediaRepository(conn *pgxpool.Pool) *MediaRepository {
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
