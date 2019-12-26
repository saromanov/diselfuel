package server

import (
	"time"
	consul "github.com/hashicorp/consul/api"
)

type Service struct {
	Name        string
	TTL         time.Duration
	ConsulAgent *consul.Agent
}

func New()(*Service, error) {
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, err
	}

	serviceDef := &consul.AgentServiceRegistration{
		Name: s.Name,
		Check: &consul.AgentServiceCheck{
			TTL: s.TTL.String(),
		},
	}

	if err := s.ConsulAgent.ServiceRegister(serviceDef); err != nil {
		return nil, err
	}
	
	return &Service {
		ConsulAgent: c,
	}
	
}