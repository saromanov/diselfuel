package app

import (
	"fmt"
	"sync"
	"time"

	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/discovery"
	"github.com/saromanov/diselfuel/internal/discovery/serf"
	"github.com/saromanov/diselfuel/internal/exec"
	"github.com/saromanov/diselfuel/internal/models"
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
		return nil, fmt.Errorf("unable to initialize service")
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
		return nil, fmt.Errorf("unable to initialize service")
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
func (a *App) Exec(query, command string) ([]*models.Exec, error) {
	nodes, err := a.serv.ListNodes(models.FilterNodes{})
	if err != nil {
		return nil, fmt.Errorf("unable to get list of nodes: %v", err)
	}

	filteredNodes, err := filterNodes(query, nodes)
	if err != nil {
		return nil, fmt.Errorf("unable to filter nodes: %v", err)
	}
	return a.exec(filteredNodes, query, command)
}

// exec provides execution of the command on target host
func (a *App) exec(hosts []*models.Host, query, command string) ([]*models.Exec, error) {
	response := []*models.Exec{}
	var wg sync.WaitGroup
	mux := &sync.Mutex{}
	wg.Add(len(hosts))
	for _, ad := range hosts {
		done := make(chan struct{})
		go func(host *models.Host) {
			defer func() {
				done <- struct{}{}
				wg.Done()
			}()
			result, err := exec.Run(command, host.Address, host.User, query)
			if err != nil {
				mux.Lock()
				response = append(response, &models.Exec{Status: models.Failed, Error: err.Error(), Host: host.Address, Name: host.Name})
				mux.Unlock()
				return
			}
			mux.Lock()
			response = append(response, &models.Exec{Status: models.Success, Output: result, Host: host.Address, Name: host.Name})
			mux.Unlock()
		}(ad)

		select {
		case <-done:
			continue
		case <-time.After(15 * time.Second):
			mux.Lock()
			response = append(response, &models.Exec{Status: models.Timeout, Host: ad.Address, Name: ad.Name})
			mux.Unlock()
			wg.Done()
		}
	}
	wg.Wait()
	return response, nil
}
