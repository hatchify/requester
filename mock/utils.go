package mock

// RequestSample represents a request sample
type RequestSample struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Body   string `json:"body,omitempty"`
}

// ResponseSample represents a response sample
type ResponseSample struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body,omitempty"`
}
