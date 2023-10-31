package cmd

import (
	"context"
	"fmt"
	"os"

	"bacakomik/config"
	"bacakomik/http"
	"bacakomik/http/chapter"
	"bacakomik/http/manga"
	"bacakomik/http/media"
	"bacakomik/repository"
	"bacakomik/repository/mysql"
	"bacakomik/service"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func RunServer(config *config.Config) {
	DSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Password, config.DBHost, config.DBPort, config.DBName)
	ctx := context.Background()
	c, err := repository.Connect(ctx, DSN)
	if err != nil {
		log.Err(err).Msg("")
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

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "start",
				Aliases:     []string{"p"},
				Description: "application running on defined port",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "port",
						Value:    8000,
						Usage:    "p",
						Aliases:  []string{"p"},
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					p := ctx.Int("port")
					ports := fmt.Sprintf(":%d", p)
					h.RunHttpServer(ports)
					return nil
				},
			},
		},
		Authors: []*cli.Author{{
			Name:  "Hanan",
			Email: "Asyrawi",
		}},
		Copyright: "readmanga",
	}

	// command run
	if err := app.Run(os.Args); err != nil {
		log.Err(err).Msg("")
	}
}
