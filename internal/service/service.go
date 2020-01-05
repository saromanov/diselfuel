package service

import (
	"fmt"
	"time"

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

	return &Service{
		ConsulClient: c,
	}, nil

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
