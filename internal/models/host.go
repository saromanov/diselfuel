package models

// Host defines representation of host
type Host struct {
	Address string   `json:"address"`
	Name    string   `json:"name"`
	Status  string   `json:"status"`
	User    string   `json:"user"`
	Tags    []string `json:"tags"`
}
