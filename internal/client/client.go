package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/saromanov/diselfuel/internal/config"
	"github.com/saromanov/diselfuel/internal/models"
)

// Client provides implementation of client
// for execution of commands
type Client struct {
	client  *http.Client
	cfg     *config.Config
	address string
}

// New provides initialization of client
func New(cfg *config.Config, address string) *Client {
	return &Client{
		client:  &http.Client{},
		cfg:     cfg,
		address: address,
	}
}

// GetList returns list of hosts
func (c *Client) GetList() ([]*models.Host, error) {
	address := c.getAddress()
	resp, err := c.client.Get(fmt.Sprintf("%s/v1/nodes", address))
	if err != nil {
		return nil, fmt.Errorf("unable to get list of hosts: %v", err)
	}

	defer resp.Body.Close()

	data := []*models.Host{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf(decodeErrTmpl, "GetList", err)
	}

	return data, nil
}

func (c *Client) getAddress() string {
	address := c.address
	fmt.Println("ADD: ", address)
	if address != "" {
		return address
	}
	if c == nil || c.cfg == nil || c.cfg.Master == nil || c.cfg.Master.Address == "" {
		panic("unable to get master host")
	}

	return c.cfg.Master.Address
}
