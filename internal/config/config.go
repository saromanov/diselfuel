package config

import (
	"fmt"

	"github.com/saromanov/cowrow"
)

// Config provides definition of configuration
type Config struct {
	Master *Server  `yaml:"master"`
	Slaves []Server `yaml:"slaves"`
	Tags   []string `yaml:"tags"`
}

// Server provides definition of the server for configuration
// It should contains Name and Addrss of the server
type Server struct {
	Name             string   `yaml:"name"`
	Address          string   `yaml:"address"`
	Port             int      `yaml:"port"`
	DiscoveryAddress string   `yaml:"discovery_address"`
	DiscoveryPort    int      `yaml:"discovery_port"`
	Tags             []string `yaml:"tags"`
	User             string   `yaml:"user"`
}

// makeDefault filling default attributes at the config
func (c *Config) makeDefault() {
	if c == nil {
		c = &Config{}
	}
	if c.Master == nil {
		c.Master = &Server{
			Name:             "default",
			Address:          "127.0.0.1",
			Port:             8080,
			DiscoveryAddress: "127.0.0.1",
		}
	}
}

// Load provides loading of the config
func Load(path string) (*Config, error) {

	c := &Config{}
	if err := cowrow.LoadByPath(path, &c); err != nil {
		return nil, fmt.Errorf("unable to load config: %v", err)
	}

	c.makeDefault()
	return c, nil
}
