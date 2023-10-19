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
func (m *MangaRepository) Create(ctx context.Context, manga *entity.Manga) {
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
	}
	log.Info().Msg("[mysql](Create): Success create")
}

// FindById method
func (m *MangaRepository) FindById(ctx context.Context, id int) {
	var manga entity.Manga
	sqlString := `SELECT * FROM mangas where id = $1`
	if err := pgxscan.Get(ctx, m.conn, &manga, sqlString, id); err != nil {
		log.Err(err).Msg("[mysql](FindById)")
	}
}
