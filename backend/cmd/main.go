package main

import (
	"bacakomik/http"
)

func main() {
	h := http.NewHTTPServer()
	// Register The Routes
	mangaHttp := http.NewMangaHttpServer(h)
	http.RegisterHttp(mangaHttp)
	h.RunHttpServer(":3000")
}
