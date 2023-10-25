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
func NewChapterRepository(conn *pgx.Conn) *ChapterRepositry {
	return &ChapterRepositry{
		conn: conn,
	}
}

// Create Data
// Create Chapter by accept manga_id
func (ch *ChapterRepositry) Create(ctx context.Context, data *entity.Chapter) (int, error) {
	var chapterID int
	// Insert into chapter
	sqlString := `INSERT INTO chapters (mangas_id, chapter , content)
					VALUES ($1 , $2 ,$3) returning id`

	r := ch.conn.QueryRow(ctx, sqlString, data.MangaID, data.Chapter, data.Content)
	if err := r.Scan(&chapterID); err != nil {
		log.Err(err).Msg("")
		return 0, err
	}
	return chapterID, nil
}

// Get All Data
func (ch *ChapterRepositry) GetAll(ctx context.Context) []*entity.Chapter {
	sqlString := `select * from chapters order by chapter desc LIMIT 200;`
	var out []*entity.Chapter

	r, err := ch.conn.Query(ctx, sqlString)
	if err != nil {
		log.Err(err).Msg("")
	}

	for r.Next() {
		c := new(entity.Chapter)
		if err := r.Scan(&c.ID, &c.MangaID, &c.Chapter, &c.Content); err != nil {
			log.Err(err).Msg("[mysql](GetAll)")
		}
		out = append(out, c)
	}

	return out
}

// Retrive One Data
func (ch *ChapterRepositry) GetOne(ctx context.Context, id int) *entity.Chapter {
	c := new(entity.Chapter)
	sqlString := `select id, chapter, content from chapters as c where c.id = $1 `
	r := ch.conn.QueryRow(ctx, sqlString, id)
	if err := r.Scan(&c.ID, &c.Chapter, &c.Content); err != nil {
		log.Err(err).Msg("")
	}
	return c
}

// Update Data
func (ch *ChapterRepositry) Update(ctx context.Context, data *entity.Chapter, id int) error {
	return nil
}

// Delete the record
func (ch *ChapterRepositry) Delete(ctx context.Context, id int) bool {
	sqlString := `delete from chapters as c where c.id = $1`
	ct, err := ch.conn.Exec(ctx, sqlString, id)
	if err != nil {
		log.Err(err).Msg("[mysql](delete):")
		return false
	}
	return ct.Delete()
}

// Get Access to instance of of T
func (ch *ChapterRepositry) NewApi() *ChapterRepositry {
	return ch
}

// GetMedias method
// Get Access Into Media
func (ch *ChapterRepositry) GetMedias(ctx context.Context, model_id int) []*entity.Media {
	c := new(entity.Chapter)
	var medias []*entity.Media
	sqlString := `select id, model_type, model_id , url from medias as m where m.model_id = $1 and m.model_type = $2;`
	r, err := ch.conn.Query(ctx, sqlString, model_id, c.String())
	if err != nil {
		log.Err(err).Msg("")
	}
	for r.Next() {
		m := new(entity.Media)
		r.Scan(&m.ID, &m.ModelType, &m.ModelType, &m.URL)
		medias = append(medias, m)
	}
	return medias
}
