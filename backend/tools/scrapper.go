package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"bacakomik/tools/provider/komikcast"
)

func main() {
	app := &cli.App{
		Name: "worker",
		Commands: []*cli.Command{
			{
				Name: "worker",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "size",
						Value:    20,
						Usage:    "w",
						Aliases:  []string{"w"},
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					i := ctx.Int("size")
					komikcast.Start(i)
					return nil
				},
			},
		},
		Authors: []*cli.Author{{
			Name:  "Hanan",
			Email: "Asyrawi",
		}},
		Copyright: "colyty",
	}

	if err := app.Run(os.Args); err != nil {
		log.Err(err).Msg("[main]")
	}
}
