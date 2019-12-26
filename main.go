package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "diselfuel",
		Usage: "Starting of the app",
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"i"},
				Usage:   "starting of the server",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "exec",
				Aliases: []string{"e"},
				Usage:   "Excecution of the command",
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