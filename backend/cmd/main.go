package main

import "bacakomik/http"

func main() {
	h := http.NewHTTPServer()
	h.RunHttpServer(":3000")
}
