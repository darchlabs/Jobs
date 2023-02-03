package providers

import (
	"github.com/darchlabs/jobs/internal/api"
	"github.com/darchlabs/jobs/internal/provider"
)

type ListProvidersHandler struct {
	providers []provider.Provider
}

func NewListProvidersHandler() *ListProvidersHandler {
	// Declare providers array
	providers := make([]provider.Provider, 0)
	// Declare each one of the providers
	dlNetworks := make([]string, 0)
	chainlinkNetworks := make([]string, 0)

	// Define Darch Labs Jobs provider with its info
	dlNetworks = append(dlNetworks, "ethereum", "polygon", "goerli")
	dlKeepers := provider.Provider{
		ID:       "1",
		Name:     "Darch Labs Keepers",
		Networks: dlNetworks,
	}

	// Define Chainlink Keepers provider with its info (empty for)
	chainlinkKeepers := provider.Provider{
		ID:       "2",
		Name:     "Chainlink Keepers",
		Networks: chainlinkNetworks,
	}

	// Append them in the array
	providers = append(providers, dlKeepers, chainlinkKeepers)

	return &ListProvidersHandler{
		providers: providers,
	}
}

func (lp *ListProvidersHandler) Invoke() *api.HandlerRes {
	// Get providers
	providers := lp.providers

	// prepare response
	return &api.HandlerRes{Payload: providers, HttpStatus: 200, Err: nil}
}
