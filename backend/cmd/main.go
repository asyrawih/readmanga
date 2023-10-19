package main

import (
	"context"
	"log"
	"time"

	"bacakomik/http"
	"bacakomik/http/manga"
	"bacakomik/record/entity"
	"bacakomik/repository"
	"bacakomik/repository/mysql"
)

func main() {
	h := http.NewHTTPServer()
	// Register The Routes
	mangaHttp := manga.NewMangaHttpServer(h)
	http.RegisterHttp(mangaHttp)

	ctx := context.Background()

	c, err := repository.Connect(ctx, "postgres://postgres:postgres@localhost:5432/readmanga")
	if err != nil {
		log.Fatal(err)
	}

	mr := mysql.NewMangaRepository(c)
	mr.Create(ctx, &entity.Manga{
		Title:        "Apalah",
		Status:       "ongoing",
		ReleaseDate:  "2023",
		TotalChapter: 200,
		Author:       "hanan",
		Type:         "manga",
		Sinopsis:     "Tidak ada",
		CreatedBy:    1,
		CreatedAt:    time.Now(),
	})

	h.RunHttpServer(":3000")
}
