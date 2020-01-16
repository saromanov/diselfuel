package config

// Execution defines configuration for execution of the
// pre-defined tasks
type Execution struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name string `json:"name"`
}
