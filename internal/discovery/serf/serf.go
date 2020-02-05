package serf

import (
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/hashicorp/serf/serf"
	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/discovery"
	"github.com/saromanov/diselfuel/internal/models"
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
	defConf.Tags = addTags(conf.Tags)
	defConf.MemberlistConfig.BindAddr = conf.Master.DiscoveryAddress
	defConf.MemberlistConfig.BindPort = conf.Master.DiscoveryPort
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

// NewStrict provides initialization of the Serf client
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

func addTags(t []string) map[string]string {
	tags := make(map[string]string)
	for i, n := range t {
		tags[fmt.Sprintf("tag_%d", i)] = n
	}

	return tags
}

// join provides joining of children nodes to network
func join(c *serf.Serf, conf *config.Config) error {
	if len(conf.Slaves) == 0 {
		return nil
	}

	nodes := []string{}
	for _, s := range conf.Slaves {
		log.Infof("Joining of nodes to the network: %s", s.Name)
		nodes = append(nodes, fmt.Sprintf("%s:%d", s.Address, s.Port))
	}

	if len(nodes) == 0 {
		log.Infof("Network not contains any nodes")
		return nil
	}
	if _, err := c.Join(nodes, true); err != nil {
		return fmt.Errorf("unable to join nodes: %v", err)
	}
	return nil
}

// ListNodes return list of nodes
func (s *Service) ListNodes(req models.FilterNodes) ([]*models.Host, error) {

	members := s.Client.Members()
	nodesResp := []*models.Host{}
	for _, n := range members {
		tags := convertTags(n.Tags)
		if req.Tag != "" {
			if findInTags(req.Tag, tags) {
				nodesResp = append(nodesResp, convertHost(n))
			}
			continue
		}
		if req.Name != "" {
			if req.Name == n.Name {
				nodesResp = append(nodesResp, convertHost(n))
			}
			continue
		}
		nodesResp = append(nodesResp, convertHost(n))
	}

	return nodesResp, nil

}

// convertTags provides converting of tags from serf representation
// to the slice of words
func convertTags(tags map[string]string) []string {
	result := []string{}
	for _, v := range tags {
		result = append(result, v)
	}
	return result
}

// convertHost provides converting of serf representation
// to the inner representation
func convertHost(n serf.Member) *models.Host {
	return &models.Host{
		Address: n.Addr.String(),
		Name:    n.Name,
		Status:  n.Status.String(),
		Tags:    convertTags(n.Tags),
	}
}

func findInTags(tag string, tags []string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

func (s *Service) Start() {

}
