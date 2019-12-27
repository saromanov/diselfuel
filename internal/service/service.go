package service

import (
	"fmt"
	"time"

	consul "github.com/hashicorp/consul/api"
	"github.com/saromanov/diselfuel/internal/config"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Name        string
	TTL         time.Duration
	ConsulAgent *consul.Client
}

// New provides initialization of the service
func New(servers *config.Config, log *logrus.Logger) (*Service, error) {
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, fmt.Errorf("unable to start consul client: %v", err)
	}

	serviceDef := &consul.AgentServiceRegistration{
		Address: "http://127.0.0.1:8080",
		ID:      "test",
		Name:    "test",
		Tags:    []string{"test"},
	}

	log.Info("Register of the service at Consul")
	if err := c.Agent().ServiceRegister(serviceDef); err != nil {
		return nil, fmt.Errorf("unable to register service: %v", err)
	}

	return &Service{
		ConsulAgent: c,
	}, nil

}

func (s *Service) Start() {

}
