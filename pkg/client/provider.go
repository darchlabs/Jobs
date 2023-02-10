package client

import "github.com/darchlabs/jobs/internal/provider"

type ListProvidersResponse struct {
	Data []*provider.Provider `json:"providers"`
}
