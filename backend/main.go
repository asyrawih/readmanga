package main

import (
	"context"
	"log"

	"bacakomik/http"
	"bacakomik/http/manga"
	"bacakomik/repository"
	"bacakomik/repository/mysql"
	"bacakomik/service"

	_ "bacakomik/docs"
)

//	@title			Hanan Test
//	@version		1.0.0
//	@description	Manga service api

//	@contact.name	API manga
//	@contact.email	hanan@asyrawih.id

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/
func main() {
	var DSN = "postgres://postgres:postgres@localhost:5432/readmanga"
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
