package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

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
					&cli.StringFlag{
						Name:     "config",
						Value:    "config.json",
						Usage:    "-c [path location]",
						Aliases:  []string{"c"},
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					p := ctx.Int("port")
					ports := fmt.Sprintf(":%d", p)

					s := ctx.String("config")

					c := config.NewConfig()
					c, err := c.LoadConfig(s)
					if err != nil {
						log.Err(err).Msg("")
					}
					cmd.RunServer(c, ports)
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
