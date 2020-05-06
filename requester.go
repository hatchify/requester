package requester

import (
	"bytes"
	"net/http"
	"net/url"
)

// New will create a new instance of Requester
func New(hc *http.Client, baseURL string) (rp *Requester) {
	var r Requester
	r.hc = hc
	r.baseURL = baseURL
	rp = &r
	return
}

// Requester is the management architecture for making HTTP Requests
type Requester struct {
	hc *http.Client

	baseURL string
}

// Private func that handles making http requests
func (r *Requester) request(method, path string, body []byte, opts *Opts) (resp *http.Response, err error) {
	var u *url.URL
	if u, err = getURL(r.baseURL, path); err != nil {
		return
	}

	var req *http.Request
	if req, err = http.NewRequest(method, u.String(), bytes.NewReader(body)); err != nil {
		return
	}

	if opts == nil {
		return r.hc.Do(req)
	}

	if opts.query != nil {
		u.RawQuery = opts.query.Encode()
	}

	opts.headers.ForEach(func(headerKey, headerVal string) (err error) {
		req.Header.Set(headerKey, headerVal)
		return
	})

	return r.hc.Do(req)
}

// Get will make an HTTP GET Request
func (r *Requester) Get(path string, opts *Opts) (resp *http.Response, err error) {
	return r.request(http.MethodGet, path, nil, opts)
}

// Put will make an HTTP Put Request
func (r *Requester) Put(path string, body []byte, opts *Opts) (resp *http.Response, err error) {
	return r.request(http.MethodPut, path, body, opts)
}

// Post will make an HTTP POST Request
func (r *Requester) Post(path string, body []byte, opts *Opts) (resp *http.Response, err error) {
	return r.request(http.MethodPost, path, body, opts)
}

// Delete will make an HTTP DELETE Request
func (r *Requester) Delete(path string, opts *Opts) (resp *http.Response, err error) {
	return r.request(http.MethodDelete, path, nil, opts)
}
