package client

import (
	"net/http"

	"github.com/saromanov/diselfuel/internal/config"
)

// Client provides implementation of client
// for execution of commands
type Client struct {
	client  *http.Client
	cfg     *config.Config
	execCfg *config.Execution
	address string
}

// New provides initialization of client
func New(cfg *config.Config, address string) *Client {
	return &Client{
		client:  &http.Client{},
		cfg:     cfg,
		address: address,
		execCfg: nil,
	}
}

// NewForApply provides loading of the config
// for apply command
func NewForApply(cfg *config.Config, execCfg *config.Execution, address string) *Client {
	return &Client{
		client:  &http.Client{},
		cfg:     cfg,
		address: address,
		execCfg: execCfg,
	}
}

func (c *Client) getAddress() string {
	address := c.address
	if address != "" {
		return address
	}
	if c == nil || c.cfg == nil || c.cfg.Master == nil || c.cfg.Master.Address == "" {
		panic("unable to get master host")
	}

	return c.cfg.Master.Address
}
