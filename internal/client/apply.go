package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/saromanov/diselfuel/internal/models"
)

// Apply provides sending request for execution of apply commands
func (c *Client) Apply(dataReq *models.Execution) ([]*models.Exec, error) {
	address := c.getAddress()
	path := fmt.Sprintf("%s/v1/apply", address)
	marshal, err := json.Marshal(dataReq)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal request: %v", err)
	}
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(marshal))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to send request to %s %v", path, err)
	}
	defer resp.Body.Close()

	data := []*models.Exec{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf(decodeErrTmpl, "Exec", err)
	}
	if data == nil {
		return nil, fmt.Errorf(applyErrTml, "Exec")
	}
	return data, nil
}
