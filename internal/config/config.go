package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
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
}

// makeDefault filling default attributes at the config
func (c *Config) makeDefault() {
	if c == nil {
		c = &Config{}
	}
	if c.Master == nil {
		c.Master = &Server{
			Name:    "default",
			Address: "localhost",
			Port:    8080,
		}
	}
}

// Load provides loading of the config
func Load(path string) (*Config, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to load config file: %v", err)
	}

	c := &Config{}
	err = yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal config file: %v", err)
	}

	c.makeDefault()
	return c, nil
}
