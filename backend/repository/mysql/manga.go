package mysql

import (
	"context"

	"bacakomik/record/entity"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

type MangaRepository struct {
	conn *pgx.Conn
}

// NewMangaRepository function
func NewMangaRepository(conn *pgx.Conn) *MangaRepository {
	return &MangaRepository{
		conn: conn,
	}
}

// Create method
func (m *MangaRepository) Create(ctx context.Context, manga *entity.Manga) error {
	sqlString := `
	        INSERT INTO mangas (title, status, release_date, total_chapter, author, type, sinopsis, created_by)
			       VALUES ($1, $2,$3,$4,$5,$6,$7,$8)`
	log.Info().Msgf("[mysql](Create): create manga %v", manga)

	if _, err := m.conn.Exec(
		ctx,
		sqlString,
		manga.Title,
		manga.Status,
		manga.ReleaseDate,
		manga.TotalChapter,
		manga.Author,
		manga.Type,
		manga.Sinopsis,
		manga.CreatedBy,
	); err != nil {
		log.Err(err).Msg("[mysql](Create)")
		return err
	}
	log.Info().Msg("[mysql](Create): Success create")
	return nil
}

// Update Data
func (ma *MangaRepository) Update(ctx context.Context, data *entity.Manga, id int) error {
	panic("not implemented") // TODO: Implement
}

// Get All Data
func (ma *MangaRepository) GetAll(ctx context.Context) []*entity.Manga {
	panic("not implemented") // TODO: Implement
}

// Retrive One Data
func (m *MangaRepository) GetOne(ctx context.Context, id int) *entity.Manga {
	var manga entity.Manga
	sqlString := `SELECT * FROM mangas where id = $1`
	if err := pgxscan.Get(ctx, m.conn, &manga, sqlString, id); err != nil {
		log.Err(err).Msg("[mysql](FindById)")
	}
	return nil
}

// Delete method
func (m *MangaRepository) Delete(ctx context.Context, id int) bool {
	sqlString := `delete from mangas where id = $1`
	ct, err := m.conn.Exec(ctx, sqlString, id)
	if err != nil {
		log.Err(err).Msg("[mysql](FindById)")
	}
	return ct.Delete()
}
