package config

// Define config variables for env
type Config struct {
	DatabaseURL string `envconfig:"database_filepath" required:"true"`
	Port        string `envconfig:"port" required:"true"`
	DockerPass  string `envconfig:"docker_pass" required:"false"`
	Version     string `envconfig:"version" required:"false"`
}
