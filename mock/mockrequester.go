package mock

/**
MockRequester simulates a regular Requester using provided data from the backing Store that holds all the requests
*/

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/hatchify/requester"
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
func (r *MockRequester) Request(method, path string, body []byte, opts requester.Opts) (resp *http.Response, err error) {
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

// Get will make an HTTP GET Request
func (r *MockRequester) Get(path string, opts ...requester.Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodGet, path, nil, opts)
}

// Put will make an HTTP Put Request
func (r *MockRequester) Put(path string, body []byte, opts ...requester.Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPut, path, body, opts)
}

// Patch will make an HTTP Patch Request
func (r *MockRequester) Patch(path string, body []byte, opts ...requester.Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPatch, path, body, opts)
}

// Post will make an HTTP POST Request
func (r *MockRequester) Post(path string, body []byte, opts ...requester.Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPost, path, body, opts)
}

// Delete will make an HTTP DELETE Request
func (r *MockRequester) Delete(path string, opts ...requester.Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodDelete, path, nil, opts)
}
