package app

import (
	"fmt"

	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/service"
	"github.com/sirupsen/logrus"
)

// App provides definition of the app
type App struct {
	conf    *config.Config
	serv    *service.Service
	logger  *logrus.Logger
	servers []config.Server
}

// New provides initialization of the app
func New(c *config.Config, log *logrus.Logger) (*App, error) {
	serv, err := service.New(c, log)
	if err != nil {
		return nil, fmt.Errorf("unable to initalize service")
	}

	return &App{
		conf:   c,
		serv:   serv,
		logger: log,
	}, nil
}

// Start provides initialization of the app
func (a *App) Start() error {
	a.serv.Start()
	return nil
}

// GetService return service
func (a *App) GetService() *service.Service {
	return a.serv
}

// Exec provides remote command execution
func (a *App) Exec() error {
	return nil
}
