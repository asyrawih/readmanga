package cmd

import (
	"context"
	"log"

	"bacakomik/http"
	"bacakomik/http/chapter"
	"bacakomik/http/manga"
	"bacakomik/repository"
	"bacakomik/repository/mysql"
	"bacakomik/service"
)

func RunServer() {
	var DSN = "postgres://postgres:postgres@localhost:5432/readmanga"
	ctx := context.Background()
	c, err := repository.Connect(ctx, DSN)
	if err != nil {
		log.Fatal(err)
	}

	mr := mysql.NewMangaRepository(c)
	ms := service.NewMangaService(mr)

	cr := mysql.NewChapterRepository(c)
	cs := service.NewChapterService(cr)

	h := http.NewHTTPServer()

	// Register The Routes
	mangaHttp := manga.NewMangaHttpServer(h, ms)
	chapterHttp := chapter.NewChapterHTTP(h, cs)
	http.RegisterHttp(mangaHttp, chapterHttp)

	h.RunHttpServer(":3000")
}
