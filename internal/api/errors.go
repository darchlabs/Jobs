package api

type Error string

const (
	// Cronjob check possible errors
	ErrorInvalidNetwork    Error = "INVALID_NETWORK"
	ErrorInvalidClient     Error = "INVALID_CLIENT"
	ErrorInvalidSigner     Error = "INVALID_SIGNER"
	ErrorInvalidAbi        Error = "INVALID_ABI"
	ErrorInexistentAddress Error = "INEXISTENT_ADDRESS"
	ErrorInvalidContract   Error = "INVALID_CONTRACT"
)
