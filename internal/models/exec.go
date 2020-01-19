package models

// Exec defines representation of the exec command
type Exec struct {
	Status string `json:"status"`
	Output []byte `json:"output"`
	Host   string `json:"host"`
	Name   string `json:"name"`
	Error  string `json:"error"`
}
