package cmd

import (
	"context"
	"fmt"

	"bacakomik/config"
	"bacakomik/http"
	"bacakomik/http/chapter"
	"bacakomik/http/manga"
	"bacakomik/http/media"
	"bacakomik/repository"
	"bacakomik/repository/mysql"
	"bacakomik/service"

	"github.com/rs/zerolog/log"
)

// Run Server
func RunServer(config *config.Config, ports string) {
	DSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Password, config.DBHost, config.DBPort, config.DBName)
	ctx := context.Background()
	c, err := repository.Connect(ctx, DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("")
		return
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
	mediaHTTP := media.NewMediaHTTPServer(h, meds, config)
	http.RegisterHttp(mangaHTTP, chapterHTTP, mediaHTTP)

	h.RunHttpServer(ports)
}
