package main

import (
	"context"
	"log"

	"bacakomik/http"
	"bacakomik/http/manga"
	"bacakomik/repository"
	"bacakomik/repository/mysql"
	"bacakomik/service"
)

var DSN = "postgres://postgres:postgres@localhost:5432/readmanga"

func main() {
	ctx := context.Background()
	c, err := repository.Connect(ctx, DSN)
	if err != nil {
		log.Fatal(err)
	}

	mr := mysql.NewMangaRepository(c)
	ms := service.NewMangaService(mr)

	h := http.NewHTTPServer()

	// Register The Routes
	mangaHttp := manga.NewMangaHttpServer(h, ms)
	http.RegisterHttp(mangaHttp)

	h.RunHttpServer(":3000")
}
