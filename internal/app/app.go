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

// NewService provides initialization of the app
// with initialization of service
func NewService(c *config.Config, log *logrus.Logger) (*App, error) {
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

// New provides initialization of instance
func New(c *config.Config, log *logrus.Logger) (*App, error) {
	return &App{
		conf:   c,
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
