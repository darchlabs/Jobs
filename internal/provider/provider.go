package provider

// Struct for DB
type Provider struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Networks []string `json:"networks"`
}

// Define state variables for the provider's jobs
type State string

const (
	StatusIdle        State = "idle"
	StatusRunning     State = "running"
	StatusStopped     State = "stopped"
	StatusError       State = "error"
	StatusAutoStopped State = "autoStopped"
)
