package models

// Status defines response after execution
type Status int

const (
	// Failed Status returns if unable to execute command
	Failed = iota
	// Success status return if command was executed correct
	Success
)

// String returns string representation of the status
func (st Status) String() string {
	switch st {
	case Failed:
		return "fail"
	case Success:
		return "success"
	}

	return "unknown"
}

// Exec defines representation of the exec command
type Exec struct {
	Status Status `json:"status"`
	Output []byte `json:"output"`
	Host   string `json:"host"`
	Name   string `json:"name"`
	Error  string `json:"error"`
}
