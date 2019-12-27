package app

import (
	"fmt"

	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/server"
	"github.com/saromanov/diselfuel/internal/service"
	"github.com/sirupsen/logrus"
)

// App provides definition of the app
type App struct {
	conf *config.Config
	serv *service.Service
}

// New provides initialization of the app
func New(c *config.Config, log *logrus.Logger) (*App, error) {
	serv, err := service.New(c)
	if err != nil {
		return nil, fmt.Errorf("unable to initalize service")
	}

	return &App{
		conf: c,
		serv: serv,
	}, nil
}

// Start provides initialization of the app
func (a *App) Start() error {
	a.serv.Start()
	server.New(a.conf)
	return nil
}