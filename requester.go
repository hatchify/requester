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

	r.setQuery(u, opts)

	var req *http.Request
	if req, err = http.NewRequest(method, u.String(), bytes.NewReader(body)); err != nil {
		return
	}

	r.setHeaders(req, opts)

	if err = r.modify(req, opts); err != nil {
		return
	}

	return r.hc.Do(req)
}

// Private func will modify a request prior to execution
func (r *Requester) modify(req *http.Request, opts *Opts) (err error) {
	for _, modifier := range opts.modifiers {
		if err = modifier(req); err != nil {
			return
		}
	}

	return
}

// Private func that will set the query pararms for a request
func (r *Requester) setQuery(u *url.URL, opts *Opts) {
	if opts == nil {
		return
	}

	if opts.query != nil {
		u.RawQuery = opts.query.Encode()
	}
}

// Private func that will set the headers for a request
func (r *Requester) setHeaders(req *http.Request, opts *Opts) {
	if opts == nil {
		return
	}

	opts.headers.ForEach(func(headerKey, headerVal string) (err error) {
		req.Header.Set(headerKey, headerVal)
		return
	})
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
