package api

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

// handler's response
type HandlerRes struct {
	Payload    interface{}
	HttpStatus int
	Err        error
}
