package main

import (
	"bacakomik/cmd"
	_ "bacakomik/docs"
)

//	@title			Manga service API
//	@version		1.0.0
//	@description	Manga service api

//	@contact.name	API manga
//	@contact.email	hanan@asyrawih.id

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/
func main() {
	cmd.RunServer()
}
