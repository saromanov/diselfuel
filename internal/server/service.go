package server

import (
	"time"

	consul "github.com/hashicorp/consul/api"
	"github.com/saromanov/diselfuel/internal/config"
)

type Service struct {
	Name        string
	TTL         time.Duration
	ConsulAgent *consul.Client
}

func New(servers *config.Config) (*Service, error) {
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, err
	}

	serviceDef := &consul.AgentServiceRegistration{
		Address: "http://127.0.0.1:8080",
		ID:      "test",
		Name:    "test",
		Tags:    []string{"test"},
	}

	if err := c.Agent().ServiceRegister(serviceDef); err != nil {
		return nil, err
	}

	return &Service{
		ConsulAgent: c,
	}, nil

}

func (s *Service) Start() {

}
