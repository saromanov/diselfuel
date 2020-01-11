package main

import (
	"os"

	"github.com/saromanov/diselfuel/internal/app"
	"github.com/saromanov/diselfuel/internal/client"
	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/server"
	"github.com/saromanov/tables"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func start(c *cli.Context) error {
	conf, err := config.Load("config.yaml")
	if err != nil {
		logrus.WithError(err).Fatal("unable to load config")
	}

	log := logger()
	a, err := app.NewService(conf, log)
	if err != nil {
		logrus.WithError(err).Fatal("unable to initialize app")
	}
	a.Start()
	server.New(a, conf, log)
	return nil
}

// exec provides execution of commands
func exec(c *cli.Context) error {
	conf, err := config.Load("config.yaml")
	if err != nil {
		logrus.WithError(err).Fatal("unable to load config")
	}

	log := logger()
	a, err := app.New(conf, log)
	if err != nil {
		logrus.WithError(err).Fatal("unable to initialize app")
	}
	a.Exec()
	return nil
}

// list returns list of nodes
func list(c *cli.Context) error {
	conf, err := config.Load("config.yaml")
	if err != nil {
		logrus.WithError(err).Fatal("unable to load config")
	}
	item := client.New(conf, "http://127.0.0.1:8081")
	resp, err := item.GetList()
	if err != nil {
		logrus.WithError(err).Fatal("unable to get list")
	}

	tab := tables.New()
	tab.AddHeader("address", "name", "status")
	for _, n := range resp {
		tab.AddLine(n.Address, n.Name, n.Status)
	}
	tab.Build()
	return nil
}
func logger() *logrus.Logger {
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
				Action:  exec,
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "Return list of hosts",
				Action:  list,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
