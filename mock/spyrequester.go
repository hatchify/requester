package mock

/**
SpyRequester is a pass-thru to Requester and therefore behaves exactly the same, however it saves all the requests into the Store
*/
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hatchify/requester"
)

// NewSpyRequester create an instance of mock requester
func NewSpyRequester(baseURL string, be Backend) (rp *SpyRequester, err error) {
	var r SpyRequester
	r.baseURL = baseURL
	if r.store, err = NewStore(be); err != nil {
		return
	}
	r.regRequester = requester.New(&http.Client{}, baseURL)
	rp = &r
	return
}

// SpyRequester implements mock requester struct
type SpyRequester struct {
	baseURL      string
	store        *Store
	regRequester *requester.Requester
}

// Request func that handles making http requests
func (r *SpyRequester) Request(method, path string, body []byte, opts requester.Opts) (resp *http.Response, err error) {

	var (
		reqSample RequestSample
		resSample ResponseSample
	)

	//Let's save that request
	reqSample = RequestSample{method, path, string(body)}

	//Logic from regular requester runs here
	resp, err = r.regRequester.Request(method, path, body, opts)

	//We are going to take the body and put it back to make it look like nothing ever happened
	tempBody, _ := ioutil.ReadAll(resp.Body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(tempBody))

	//Forming our ResponseSample
	resSample = ResponseSample{resp.StatusCode, string(tempBody)}

	//Let's save our request to db
	r.store.Set(reqSample, resSample)

	r.store.Save() //Let's save our db :)

	//Debugging data
	//Our request parameters
	fmt.Println(method, path, string(body))
	fmt.Println(string(tempBody))
	fmt.Println("---Saved---")

	return
}

// Get will make an HTTP GET Request
func (r *SpyRequester) Get(path string, opts ...requester.Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodGet, path, nil, opts)
}

// Put will make an HTTP Put Request
func (r *SpyRequester) Put(path string, body []byte, opts ...requester.Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPut, path, body, opts)
}

// Patch will make an HTTP Patch Request
func (r *SpyRequester) Patch(path string, body []byte, opts ...requester.Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPatch, path, body, opts)
}

// Post will make an HTTP POST Request
func (r *SpyRequester) Post(path string, body []byte, opts ...requester.Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPost, path, body, opts)
}

// Delete will make an HTTP DELETE Request
func (r *SpyRequester) Delete(path string, opts ...requester.Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodDelete, path, nil, opts)
}
