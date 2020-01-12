package client

import (
	"encoding/json"
	"fmt"

	"github.com/saromanov/diselfuel/internal/models"
)

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
