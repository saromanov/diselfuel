package serf

import (
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/hashicorp/serf/serf"
	"github.com/saromanov/diselfuel/internal/config"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Name         string
	TTL          time.Duration
	Client *serf.Client
}

// New provides initialization of the service
// with registration of the new client
func New(conf *config.Config, log *logrus.Logger) (*Service, error) {
	c, err := serf.Create(serf.DefaultConfig())
	if err != nil {
		return nil, fmt.Errorf("unable to start serf client: %v", err)
	}

	return &Service{
		Client: c,
	}, nil

}

// NewStrict provides initialization of the Consul client
func NewStrict(conf *config.Config, log *logrus.Logger) (*Service, error) {
	c, err := serf.Create(serf.DefaultConfig())
	if err != nil {
		return nil, fmt.Errorf("unable to start serf client: %v", err)
	}
	return &Service{
		Client: c,
	}, nil
}

// join provides joining of children nodes to network
func join(c *consul.Client, conf *config.Config) error {
	if len(conf.Slaves) == 0 {
		return nil
	}

	for _, s := range conf.Slaves {
		serviceDef := &consul.AgentServiceRegistration{
			Address: fmt.Sprintf("%s:%d", s.Address, s.Port),
			ID:      s.Name,
			Name:    s.Name,
			Tags:    []string{"test"},
		}

		log.Infof("Joining of nodes to the network: %s", s.Name)
		if err := c.Agent().ServiceRegister(serviceDef); err != nil {
			return fmt.Errorf("unable to register service %s: %v", s.Name, err)
		}
	}

	return nil
}

// ListNodes return list of nodes
func (s *Service) ListNodes() ([]string, error) {
	nodes, _, err := s.Client.Members()
	if err != nil {
		return nil, fmt.Errorf("unable to get list of nodes: %v", err)
	}
	nodesResp := make([]string, len(nodes))
	for i, n := range nodes {
		nodesResp[i] = n.Addr.String()
	}

	return nodesResp, nil
}
func (s *Service) Start() {

}
