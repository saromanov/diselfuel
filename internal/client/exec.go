package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/saromanov/diselfuel/internal/models"
)

// Exec returns list of hosts
func (c *Client) Exec(query, command string) ([]*models.Exec, error) {
	address := c.getAddress()
	path := fmt.Sprintf("%s/v1/exec", address)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %v", err)
	}

	q := req.URL.Query()
	q.Add("query", query)
	q.Add("command", command)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to send request to %s %v", path, err)
	}
	defer resp.Body.Close()

	data := []*models.Exec{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf(decodeErrTmpl, "Exec", err)
	}
	return data, nil
}
