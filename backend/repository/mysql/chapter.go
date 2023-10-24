package mysql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"bacakomik/record/entity"
)

type ChapterRepositry struct {
	conn *pgx.Conn
}

// NewMangaRepository function
func NewChapterRepository(conn *pgx.Conn) *MangaRepository {
	return &MangaRepository{
		conn: conn,
	}
}

// Create Data
// Create Chapter by accept manga_id
func (ch *ChapterRepositry) Create(ctx context.Context, data *entity.Chapter) (int, error) {
	var chapterID int
	// Insert into chapter
	sqlString := `INSERT INTO chapters (manga_id, chapter , content)
					VALUES ($1 , $2 ,$2)`

	r := ch.conn.QueryRow(ctx, sqlString, data.MangaID, data.Chapter, data.Content)
	if err := r.Scan(&chapterID); err != nil {
		log.Err(err).Msg("")
		return 0, err
	}
	return chapterID, nil
}
