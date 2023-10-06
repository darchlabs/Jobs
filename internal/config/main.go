package config

// Define config variables for env
type Config struct {
	DatabaseFilepath string `envconfig:"database_filepath" required:"true"`
	BackofficeApiURL string `envconfig:"backoffice_api_url" required:"true"`
	Port             string `envconfig:"port" required:"true"`
}
