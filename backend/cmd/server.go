package cmd

import (
	"context"
	"log"

	"bacakomik/http"
	"bacakomik/http/chapter"
	"bacakomik/http/manga"
	"bacakomik/http/media"
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

	medr := mysql.NewMediaRepository(c)
	meds := service.NewMediaService(medr)

	// BOOTSTRAP HTTP SERVER
	h := http.NewHTTPServer()

	// Register The Routes
	mangaHTTP := manga.NewMangaHttpServer(h, ms)
	chapterHTTP := chapter.NewChapterHTTP(h, cs)
	mediaHTTP := media.NewMediaHTTPServer(h, meds)
	http.RegisterHttp(mangaHTTP, chapterHTTP, mediaHTTP)

	h.RunHttpServer(":8000")
}
