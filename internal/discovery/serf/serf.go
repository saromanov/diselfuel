package serf

import (
	"fmt"
	"time"

	"github.com/hashicorp/serf/serf"
	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/discovery"
	"github.com/saromanov/diselfuel/internal/models"
	"github.com/sirupsen/logrus"
)

const defaultAddress = "127.0.0.1"

// Service adds definition for serf discovery
type Service struct {
	Name   string
	TTL    time.Duration
	Client *serf.Serf
}

// New provides initialization of master
func New(conf *config.Config, log *logrus.Logger) (discovery.Discovery, error) {
	c, client, err := newService(conf, log)
	if _, err := client.Join(conf.Slaves, true); err != nil {
		return nil, err
	}
	return c, err
}

// general method for initialization of service
func newService(conf *config.Config, log *logrus.Logger) (discovery.Discovery, *serf.Serf, error) {
	defConf := serf.DefaultConfig()
	defConf.Tags = addTags(conf.Tags)
	defConf.MemberlistConfig.BindAddr = conf.Server.DiscoveryAddress
	defConf.MemberlistConfig.BindPort = conf.Server.DiscoveryPort
	c, err := serf.Create(defConf)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to start serf client: %v", err)
	}

	return &Service{
		Client: c,
	}, c, nil
}

// NewStrict provides initialization of the Serf client
func NewStrict(conf *config.Config, log *logrus.Logger) (*Service, error) {
	if conf == nil {
		return nil, fmt.Errorf("discovery: config is not defined")
	}
	defConf := serf.DefaultConfig()
	defConf.MemberlistConfig.AdvertiseAddr = defaultAddress
	defConf.MemberlistConfig.AdvertisePort = conf.Server.DiscoveryPort
	defConf.MemberlistConfig.BindAddr = defaultAddress
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
	if _, err := c.Join([]string{fmt.Sprintf("%s:%d", conf.Server.Address, conf.Server.Port)}, true); err != nil {
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
		if req.Name == "" && req.Tag == "" {
			nodesResp = append(nodesResp, convertHost(n))
		}
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
