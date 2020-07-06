package requester

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MockRequester implements mock requester struct
type MockRequester struct {
	baseURL string
	hc      *http.Client
	store   Store
}

// NewMock create an instance of mock requester
func NewMock(hc *http.Client, baseURL string, store Store) (rp *MockRequester) {
	var r MockRequester
	r.hc = hc
	r.baseURL = baseURL
	r.store = store
	rp = &r
	return
}

// Request func that handles making http requests
func (r *MockRequester) Request(method, path string, body []byte, opts Opts) (resp *http.Response, err error) {
	//Implement the sauce
	var (
		reqSample RequestSample
		resSample ResponseSample
	)

	//Let's save that request
	reqSample = RequestSample{method, path, string(body)}

	//So let's try mocking stuff by using only data in our db
	if resSample, err = r.store.Get(reqSample); err != nil {
		resSample = ResponseSample{
			StatusCode: 404,
			Body:       err.Error(),
		}
	}

	resp = &http.Response{
		StatusCode: resSample.StatusCode,
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(resSample.Body))),
	}
	return
}

func (r *MockRequester) setOpts(req *http.Request, opts Opts) (err error) {
	for _, opt := range opts {
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
func (r *MockRequester) setQuery(req *http.Request, query Query) {
	req.URL.RawQuery = query.Encode()
}

// Private func that will set the headers for a request, will not error
func (r *MockRequester) setHeaders(req *http.Request, headers Headers) {
	headers.ForEach(func(headerKey, headerVal string) (err error) {
		req.Header.Set(headerKey, headerVal)
		return
	})
}

// Get will make an HTTP GET Request
func (r *MockRequester) Get(path string, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodGet, path, nil, opts)
}

// Put will make an HTTP Put Request
func (r *MockRequester) Put(path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPut, path, body, opts)
}

// Patch will make an HTTP Patch Request
func (r *MockRequester) Patch(path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPatch, path, body, opts)
}

// Post will make an HTTP POST Request
func (r *MockRequester) Post(path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPost, path, body, opts)
}

// Delete will make an HTTP DELETE Request
func (r *MockRequester) Delete(path string, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodDelete, path, nil, opts)
}
