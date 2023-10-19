package mysql

import (
	"github.com/jackc/pgx/v5"
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
