package main

import (
	"github.com/rs/zerolog/log"

	"bacakomik/cmd"
	"bacakomik/config"
	_ "bacakomik/docs"
)

//	@title			Manga service API
//	@version		1.0.0
//	@description	Manga service api

//	@contact.name	API manga
//	@contact.email	hanan@asyrawih.id

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8000
// @BasePath	/
func main() {
	c := config.NewConfig()
	c, err := c.LoadConfig("config.json")
	if err != nil {
		log.Err(err).Msg("")
	}
	cmd.RunServer(c)
}
