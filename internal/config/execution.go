package config

import (
	"fmt"

	"github.com/saromanov/cowrow"
)

// Execution defines configuration for execution of the
// pre-defined tasks
type Execution struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name string `json:"name"`
}

// LoadExecConfig provides loading of the execution config
func LoadExecConfig(path string) (*Execution, error) {

	c := &Execution{}
	if err := cowrow.LoadByPath(path, &c); err != nil {
		return nil, fmt.Errorf("unable to load execution config: %v", err)
	}

	return c, nil
}
