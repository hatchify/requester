package mock

// RequestSample represents a request sample
type RequestSample struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Body   string `json:"request_body"`
}

// ResponseSample represents a response sample
type ResponseSample struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"response_body"`
}
