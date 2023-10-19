package mysql

import "net"

type MangaRepository struct {
	conn *net.Conn
}
