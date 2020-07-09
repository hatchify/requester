package requester

/**
SpyRequester is a pass-thru to Requester and therefore behaves exactly the same, however it saves all the requests into the Store
*/
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SpyRequester implements mock requester struct
type SpyRequester struct {
	baseURL      string
	hc           *http.Client
	store        Store
	regRequester *Requester
}

// NewMock create an instance of mock requester
func NewSpy(hc *http.Client, baseURL string, store Store) (rp *SpyRequester) {
	var r SpyRequester
	r.hc = hc
	r.baseURL = baseURL
	r.store = store
	r.regRequester = New(&http.Client{}, baseURL)
	rp = &r
	return
}

// Request func that handles making http requests
func (r *SpyRequester) Request(method, path string, body []byte, opts Opts) (resp *http.Response, err error) {

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

//TODO: Implement this - so that we would store all the options in our RequestSample
func (r *SpyRequester) setOpts(req *http.Request, opts Opts) (err error) {
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
func (r *SpyRequester) setQuery(req *http.Request, query Query) {
	req.URL.RawQuery = query.Encode()
}

// Private func that will set the headers for a request, will not error
func (r *SpyRequester) setHeaders(req *http.Request, headers Headers) {
	_ = headers.ForEach(func(headerKey, headerVal string) (err error) {
		req.Header.Set(headerKey, headerVal)
		return
	})
}

// Get will make an HTTP GET Request
func (r *SpyRequester) Get(path string, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodGet, path, nil, opts)
}

// Put will make an HTTP Put Request
func (r *SpyRequester) Put(path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPut, path, body, opts)
}

// Patch will make an HTTP Patch Request
func (r *SpyRequester) Patch(path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPatch, path, body, opts)
}

// Post will make an HTTP POST Request
func (r *SpyRequester) Post(path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPost, path, body, opts)
}

// Delete will make an HTTP DELETE Request
func (r *SpyRequester) Delete(path string, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodDelete, path, nil, opts)
}
