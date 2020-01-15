package app

import (
	"fmt"
	"regexp"
	"strings"

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
func (a *App) Exec(query, command string) (*models.Exec, error) {
	nodes, err := a.serv.ListNodes()
	if err != nil {
		return nil, fmt.Errorf("unable to get list of nodes: %v", err)
	}

	filteredNodes, err := filterNodes(query, nodes)
	if err != nil {
		return nil, fmt.Errorf("unable to filter nodes: %v", err)
	}
	addresses := getNodeAddresess(filteredNodes)
	fmt.Println("NOdes: ", addresses)
	for _, ad := range addresses {
		if err := exec.Run(command, ad, query, query); err != nil {
			fmt.Println("ERR: ", err)
			return &models.Exec{Status: "fail"}, err
		}
	}
	return &models.Exec{Status: "ok"}, nil
}

func getNodeAddresess(hosts []*models.Host) []string {
	response := make([]string, len(hosts))
	for i, h := range hosts {
		response[i] = h.Address
	}
	return response
}

// filterNodes provides filtering of nodes by the query
// for example:
// node=test1
// node=test*
func filterNodes(query string, hosts []*models.Host) ([]*models.Host, error) {
	if query == "*" {
		return hosts, nil
	}
	if strings.HasPrefix(query, "node") {
		rawResult := strings.Split(query, "=")
		if len(rawResult) < 2 {
			return nil, fmt.Errorf("invalid expression")
		}
		result := []*models.Host{}
		re := regexp.MustCompile(rawResult[1])
		for _, h := range hosts {
			if s := re.FindString(h.Name); s != "" {
				result = append(result, h)
			}
		}

		return result, nil
	}

	return nil, nil
}
