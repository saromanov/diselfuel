package main

import (
	"os"

	"github.com/saromanov/antenna/command/client"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "diselfuel",
		Usage: "make an explosive entrance",
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"i"},
				Usage:   "starting of the server",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}