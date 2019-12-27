package main

import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/app"
	"github.com/urfave/cli/v2"
)

func start(c *cli.Context) error {
	conf, err := config.Load("config.yaml")
	if err != nil {
		logrus.WithError(err).Fatal("unable to load config")
	}
	
	a, err := app.New(conf, logger())
	if err != nil {
		logrus.WithError(err).Fatal("unable to initialize app")
	}

	a.Start()
	return nil
}

func logger()*logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	return log
}
func main() {
	app := &cli.App{
		Name:  "diselfuel",
		Usage: "Starting of the app",
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"i"},
				Usage:   "starting of the server",
				Action:  start,
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
