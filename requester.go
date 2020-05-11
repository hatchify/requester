package requester

import (
	"bytes"
	"fmt"
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
func (r *Requester) request(method, path string, body []byte, opts Opts) (resp *http.Response, err error) {
	var u *url.URL
	if u, err = getURL(r.baseURL, path); err != nil {
		return
	}

	var req *http.Request
	if req, err = http.NewRequest(method, u.String(), bytes.NewReader(body)); err != nil {
		return
	}

	if err = r.setOpts(req, opts); err != nil {
		return
	}

	fmt.Printf("headers: %v\n\n", req.Header)
	return r.hc.Do(req)
}

func (r *Requester) setOpts(req *http.Request, opts Opts) (err error) {

	for _, opt := range opts {
		fmt.Printf("%T\n\n", opt)
		switch t := opt.(type) {
		case Query:
			r.setQuery(req, t)
		case Headers:
			r.setHeaders(req, t)
		case Modifier:
			err = t(req, r.hc)
		default:
			err = fmt.Errorf("invalid opts type: expected \"Query\", \"Headers\", or \"Modifier\", received \"%T\"", opt)
		}
	}

	return
}

// Private func that will set the query pararms for a request
func (r *Requester) setQuery(req *http.Request, query Query) {
	req.URL.RawQuery = query.Encode()
}

// Private func that will set the headers for a request, will not error
func (r *Requester) setHeaders(req *http.Request, headers Headers) {
	headers.ForEach(func(headerKey, headerVal string) (err error) {
		req.Header.Set(headerKey, headerVal)
		return
	})
}

// Get will make an HTTP GET Request
func (r *Requester) Get(path string, opts ...Opt) (resp *http.Response, err error) {
	return r.request(http.MethodGet, path, nil, opts)
}

// Put will make an HTTP Put Request
func (r *Requester) Put(path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.request(http.MethodPut, path, body, opts)
}

// Post will make an HTTP POST Request
func (r *Requester) Post(path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.request(http.MethodPost, path, body, opts)
}

// Delete will make an HTTP DELETE Request
func (r *Requester) Delete(path string, opts ...Opt) (resp *http.Response, err error) {
	return r.request(http.MethodDelete, path, nil, opts)
}
