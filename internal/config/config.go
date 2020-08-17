package config

import (
	"fmt"
	"time"

	"github.com/saromanov/cowrow"
)

// Config provides definition of configuration
type Config struct {
	Server *Server  `yaml:"server"`
	Slaves []Server `yaml:"slaves"`
	Tags   []string `yaml:"tags"`
}

// Server provides definition of the server for configuration
// It should contains Name and Addrss of the server
type Server struct {
	Name             string        `yaml:"name"`
	Address          string        `yaml:"address"`
	Port             int           `yaml:"port"`
	Type             string        `yaml:"type"`
	DiscoveryAddress string        `yaml:"discovery_address"`
	DiscoveryPort    int           `yaml:"discovery_port"`
	Tags             []string      `yaml:"tags"`
	User             string        `yaml:"user"`
	RootDir          string        `yaml:"root_dir"`
	ModuleDirs       []string      `yaml:"module_dirs"`
	CacheDir         string        `yaml:"cache_dir"`
	CommandTimeout   string        `yaml:"command_timeout"`
	CacheTimeout     time.Duration `yaml:"cache_timeout"`
}

// makeDefault filling default attributes at the config
func (c *Config) makeDefault() {
	if c == nil {
		c = &Config{}
	}
	if c.Server == nil {
		c.Server = &Server{
			Name:             "default",
			Address:          "127.0.0.1",
			Port:             8080,
			DiscoveryAddress: "127.0.0.1",
			CacheTimeout:     10 * time.Minute,
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
