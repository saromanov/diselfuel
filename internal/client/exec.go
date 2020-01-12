package client

import (
	"encoding/json"
	"fmt"

	"github.com/saromanov/diselfuel/internal/models"
)

// Exec returns list of hosts
func (c *Client) Exec(query, command string) error {
	address := c.getAddress()
	resp, err := c.client.Get(fmt.Sprintf("%s/v1/exec", address))
	if err != nil {
		return fmt.Errorf("unable to get list of hosts: %v", err)
	}

	defer resp.Body.Close()

	data := []*models.Exec{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf(decodeErrTmpl, "GetList", err)
	}

	return nil
}
