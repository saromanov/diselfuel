package config

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// Config provides definition of configuration
type Config struct {
	Sefvers []Server `yaml:"servers"`
}

// Server provides definition of the server for configuration
// It should contains Name and Addrss of the server
type Server struct {
	Name string `yaml:"name"`
	Address string `yaml:"address"`
}

// Load provides loading of the config
func Load(path string)(*Config, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
    	return nil, fmt.Errorf("unable to load config file: %v", err)
	}

	c := &Config{}
	err = yaml.Unmarshal([]byte(data), &c)
	if err != nil {
			return nil, fmt.Errorf("unable to unmarshal config file: %v", err)
	}
	return c, nil
}