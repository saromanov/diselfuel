package app

import (
	"fmt"

	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/discovery"
	"github.com/saromanov/diselfuel/internal/discovery/serf"
	"github.com/sirupsen/logrus"
)

// App provides definition of the app
type App struct {
	conf    *config.Config
	serv    discovery.Discovery
	logger  *logrus.Logger
	servers []config.Server
}

// NewService provides initialization of the app
// with initialization of service
func NewService(c *config.Config, log *logrus.Logger) (*App, error) {
	serv, err := serf.New(c, log)
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
	serv, err := serf.NewStrict(c, log)
	if err != nil {
		return nil, fmt.Errorf("unable to initalize service")
	}
	return &App{
		conf:   c,
		logger: log,
		serv:   serv,
	}, nil
}

// GetService return service
func (a *App) GetService() discovery.Discovery {
	return a.serv
}

// Exec provides remote command execution
func (a *App) Exec(query, command string) error {
	nodes, err := a.serv.ListNodes()
	if err != nil {
		return fmt.Errorf("unable to get list of nodes: %v", err)
	}

	fmt.Println("NOdes: ", nodes)
	return nil
}
