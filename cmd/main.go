package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/saromanov/diselfuel/internal/app"
	"github.com/saromanov/diselfuel/internal/client"
	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/models"
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
	server.New(a, conf, log)
	return nil
}

// exec provides execution of commands
func exec(c *cli.Context) error {
	args := c.Args()
	if args.Len() < 2 {
		logrus.Fatal("not enough arguments")
	}
	addressFlag := c.String("address")
	conf, err := config.Load("config.yaml")
	if err != nil {
		logrus.WithError(err).Fatal("unable to load config")
	}

	address := "http://127.0.0.1:8081"
	if addressFlag != "" {
		address = addressFlag
	}
	item := client.New(conf, address)
	result, err := item.Exec(args.Get(0), args.Get(1))
	if err != nil {
		logrus.WithError(err).Fatal("unable to execute command")
	}
	for i, r := range result {
		num := i + 1
		switch r.Status {
		case models.Failed:
			color.Red("%d. %s %s %s\n\n", num, r.Name, r.Host, r.Status)
		case models.Success:
			color.Green("%d. %s %s %s\n", num, r.Name, r.Host, r.Status)
			fmt.Println(string(r.Output))
		case models.Timeout:
			color.Yellow("%d. %s %s %s\n", num, r.Name, r.Host, r.Status)
		}
	}
	return nil
}

// list returns list of nodes
func list(c *cli.Context) error {
	addressFlag := c.String("address")
	conf, err := config.Load("config.yaml")
	if err != nil {
		logrus.WithError(err).Fatal("unable to load config")
	}
	address := "http://127.0.0.1:8081"
	if addressFlag != "" {
		address = addressFlag
	}
	item := client.New(conf, address)
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

func apply(c *cli.Context) error {
	addressFlag := c.String("address")
	conf, err := config.Load("config.yaml")
	if err != nil {
		logrus.WithError(err).Fatal("unable to load config")
	}
	execConfig := c.String("exec-config")
	if execConfig == "" {
		logrus.Fatal("exec config is not defined")
	}
	execConf, err := config.LoadExecConfig(execConfig)
	if err != nil {
		logrus.WithError(err).Fatalf("unable to load exec config: %v", err)
	}
	address := "http://127.0.0.1:8081"
	if addressFlag != "" {
		address = addressFlag
	}
	item := client.New(conf, address)
	_, err = item.Apply(&models.Execution{
		Tasks: convertTasks(execConf.Tasks),
	})
	if err != nil {
		logrus.WithError(err).Fatalf("unable to make Apply: %v", err)
	}
	fmt.Println(conf)
	return nil
}

func convertTasks(tasks []config.Task) []models.Task {
	resp := []models.Task{}
	for _, t := range tasks {
		resp = append(resp, models.Task{
			Name:    t.Name,
			Tag:     t.Tag,
			Command: t.Command,
		})
	}
	return resp
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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "address",
				Value: "",
				Usage: "address of the server for execution",
			},
			&cli.StringFlag{
				Name:  "config",
				Value: "config.yaml",
				Usage: "path to config",
			},
			&cli.StringFlag{
				Name:  "exec-config",
				Usage: "path to execution config",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "start",
				Usage:  "starting of the server",
				Action: start,
			},
			{
				Name:   "exec",
				Usage:  "Execution of the command",
				Action: exec,
			},
			{
				Name:   "list",
				Usage:  "Return list of hosts",
				Action: list,
			},
			{
				Name:   "apply",
				Usage:  "Applying of execution for target servers",
				Action: apply,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
