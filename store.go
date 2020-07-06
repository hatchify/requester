package requester

// Store implements storage for requests
type Store interface {
	Get(request RequestSample) (response ResponseSample, err error)
	Set(request RequestSample, response ResponseSample)
	GetAll() interface{}
	Save()
}

type RequestSample struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Body   string `json:"request_body"`
}

type ResponseSample struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"response_body"`
}
