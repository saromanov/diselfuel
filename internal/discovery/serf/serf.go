package serf

import (
	"fmt"
	"time"

	"github.com/apex/log"
	consul "github.com/hashicorp/consul/api"
	"github.com/saromanov/diselfuel/internal/config"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Name         string
	TTL          time.Duration
	ConsulClient *consul.Client
}

// New provides initialization of the service
// with registration of the new client
func New(conf *config.Config, log *logrus.Logger) (*Service, error) {
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, fmt.Errorf("unable to start consul client: %v", err)
	}

	serviceDef := &consul.AgentServiceRegistration{
		Address: fmt.Sprintf("%s:%d", conf.Master.Address, conf.Master.Port),
		ID:      conf.Master.Name,
		Name:    conf.Master.Name,
		Tags:    []string{"test"},
	}

	log.Info("Register of the service at Consul")
	if err := c.Agent().ServiceRegister(serviceDef); err != nil {
		return nil, fmt.Errorf("unable to register service: %v", err)
	}

	if err := join(c, conf); err != nil {
		return nil, fmt.Errorf("unable to join nodes: %v", err)
	}

	return &Service{
		ConsulClient: c,
	}, nil

}

// NewStrict provides initialization of the Consul client
func NewStrict(conf *config.Config, log *logrus.Logger) (*Service, error) {
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, fmt.Errorf("unable to start consul client: %v", err)
	}
	return &Service{
		ConsulClient: c,
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
	nodes, _, err := s.ConsulClient.Catalog().Nodes(&consul.QueryOptions{})
	if err != nil {
		return nil, fmt.Errorf("unable to get list of nodes: %v", err)
	}
	nodesResp := make([]string, len(nodes))
	for i, n := range nodes {
		nodesResp[i] = n.Address
	}

	return nodesResp, nil
}
func (s *Service) Start() {

}
