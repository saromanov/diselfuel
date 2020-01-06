package serf

import (
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/hashicorp/serf/serf"
	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/discovery"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Name   string
	TTL    time.Duration
	Client *serf.Serf
}

// New provides initialization of the service
// with registration of the new client
func New(conf *config.Config, log *logrus.Logger) (discovery.Discovery, error) {
	defConf := serf.DefaultConfig()
	defConf.MemberlistConfig.BindPort = 7779
	c, err := serf.Create(defConf)
	if err != nil {
		return nil, fmt.Errorf("unable to start serf client: %v", err)
	}

	if err := join(c, conf); err != nil {
		return nil, err
	}
	return &Service{
		Client: c,
	}, nil

}

// NewStrict provides initialization of the Consul client
func NewStrict(conf *config.Config, log *logrus.Logger) (*Service, error) {
	defConf := serf.DefaultConfig()
	defConf.MemberlistConfig.AdvertiseAddr = "127.0.0.1"
	defConf.MemberlistConfig.AdvertisePort = 7779
	defConf.MemberlistConfig.BindAddr = "127.0.0.1"
	defConf.MemberlistConfig.BindPort = 7781
	c, err := serf.Create(defConf)
	if err != nil {
		return nil, fmt.Errorf("unable to start serf client: %v", err)
	}

	return &Service{
		Client: c,
	}, nil
}

// join provides joining of children nodes to network
func join(c *serf.Serf, conf *config.Config) error {
	if len(conf.Slaves) == 0 {
		return nil
	}

	nodes := []string{}
	for _, s := range conf.Slaves {
		log.Infof("Joining of nodes to the network: %s", s.Name)
		nodes = append(nodes, fmt.Sprintf("%s:%d", "172.17.0.3", 7946))
	}

	if _, err := c.Join(nodes, true); err != nil {
		return fmt.Errorf("unable to join nodes: %v", err)
	}
	return nil
}

// ListNodes return list of nodes
func (s *Service) ListNodes() ([]string, error) {
	members := s.Client.Members()
	nodesResp := make([]string, len(members))
	for i, n := range members {
		nodesResp[i] = n.Addr.String()
	}

	fmt.Println("MEMBERS: ", members)
	return nodesResp, nil
}
func (s *Service) Start() {

}
