package app

import (
	"fmt"

	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/service"
)

type App struct {
	conf *config.Config
	serv *service.Service
}

// New provides initialization of the app
func New(c *config.Config) (*App, error) {
	serv, err := service.New(c)
	if err != nil {
		return nil, fmt.Errorf("unable to initalize service")
	}

	return &App{
		conf: c,
		serv: serv,
	}
}

// Start provides initialization of the app
func (a *App) Start() error {
	a.serv.Start()
	return nil
}
