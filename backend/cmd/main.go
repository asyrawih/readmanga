package main

import (
	"bacakomik/http"
	"bacakomik/http/manga"
)

func main() {
	h := http.NewHTTPServer()
	// Register The Routes
	mangaHttp := manga.NewMangaHttpServer(h)
	http.RegisterHttp(mangaHttp)
	h.RunHttpServer(":3000")
}
