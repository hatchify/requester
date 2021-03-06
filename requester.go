package requester

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
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

	// Prepender is called before each request
	prepender func() Opts
}

// Request func that handles making http requests
func (r *Requester) Request(method, path string, body []byte, opts Opts) (resp *http.Response, err error) {
	return r.RequestWithContext(context.Background(), method, path, body, opts)
}

// RequestWithContext func that handles making http requests with a given context
func (r *Requester) RequestWithContext(ctx context.Context, method, path string, body []byte, opts Opts) (resp *http.Response, err error) {
	var req *http.Request
	// Create new request
	if req, err = r.newRequest(ctx, method, path, body); err != nil {
		return
	}

	// Set request options
	if err = r.setOpts(req, opts); err != nil {
		return
	}

	// Perform request using HTTP client
	return r.hc.Do(req)
}

// Get will make an HTTP GET Request
func (r *Requester) Get(path string, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodGet, path, nil, opts)
}

// GetWithContext will make an HTTP GET Request with a given context
func (r *Requester) GetWithContext(ctx context.Context, path string, opts ...Opt) (resp *http.Response, err error) {
	return r.RequestWithContext(ctx, http.MethodGet, path, nil, opts)
}

// Post will make an HTTP POST Request
func (r *Requester) Post(path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPost, path, body, opts)
}

// PostWithContext will make an HTTP POST Request with a given context
func (r *Requester) PostWithContext(ctx context.Context, path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.RequestWithContext(ctx, http.MethodPost, path, body, opts)
}

// Put will make an HTTP Put Request
func (r *Requester) Put(path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPut, path, body, opts)
}

// PutWithContext will make an HTTP Put Request with a given context
func (r *Requester) PutWithContext(ctx context.Context, path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.RequestWithContext(ctx, http.MethodPut, path, body, opts)
}

// Patch will make an HTTP Patch Request
func (r *Requester) Patch(path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodPatch, path, body, opts)
}

// PatchWithContext will make an HTTP Patch Request with a given context
func (r *Requester) PatchWithContext(ctx context.Context, path string, body []byte, opts ...Opt) (resp *http.Response, err error) {
	return r.RequestWithContext(ctx, http.MethodPatch, path, body, opts)
}

// Delete will make an HTTP DELETE Request
func (r *Requester) Delete(path string, opts ...Opt) (resp *http.Response, err error) {
	return r.Request(http.MethodDelete, path, nil, opts)
}

// DeleteWithContext will make an HTTP DELETE Request with a given context
func (r *Requester) DeleteWithContext(ctx context.Context, path string, opts ...Opt) (resp *http.Response, err error) {
	return r.RequestWithContext(ctx, http.MethodDelete, path, nil, opts)
}

// SetOptsPrepender will set an opts prepender function for the given instance of Requester
// Note: The prepender func allows for opts to be set for all requests. This
// can be quiet useful for things like Authorization tokens
func (r *Requester) SetOptsPrepender(prepender func() Opts) {
	r.prepender = prepender
}

// GetJar will return the underlying jar of a requester instance
func (r *Requester) GetJar() (jar http.CookieJar) {
	return r.hc.Jar
}

// RequestWithContext func that handles making http requests with a given context
func (r *Requester) newRequest(ctx context.Context, method, path string, body []byte) (req *http.Request, err error) {
	var u *url.URL
	// Get URL from base URL and provided path
	if u, err = getURL(r.baseURL, path); err != nil {
		return
	}

	// Create a new HTTP request
	if req, err = http.NewRequest(method, u.String(), bytes.NewReader(body)); err != nil {
		return
	}

	// Set context for request
	req.WithContext(ctx)
	return
}

func (r *Requester) setOpts(req *http.Request, opts Opts) (err error) {
	if r.prepender != nil {
		// Prepender exists, prepend opts to opts list
		opts = append(r.prepender(), opts...)
	}

	for _, opt := range opts {
		switch t := opt.(type) {
		case Query:
			r.setQuery(req, t)
		case Header:
			r.setHeader(req, opts, t)
		case Headers:
			r.setHeaders(req, t)
		case Body:
			r.setBody(req, t)
		case BasicAuth:
			r.setBasicAuth(req, t)
		case Modifier:
			err = t(req, r.hc)
		default:
			err = fmt.Errorf("invalid type for opts: \"%T\"", opt)
		}
	}

	return
}

// Private func that will set the query pararms for a request
func (r *Requester) setQuery(req *http.Request, query Query) {
	req.URL.RawQuery = query.Encode()
}

// Private func that will set the headers for a request, will not error
func (r *Requester) setHeader(req *http.Request, opts Opts, header Header) {
	req.Header.Set(header.Key, header.Val)
}

// Private func that will set the headers for a request, will not error
func (r *Requester) setHeaders(req *http.Request, headers Headers) {
	_ = headers.ForEach(func(headerKey, headerVal string) (err error) {
		req.Header.Set(headerKey, headerVal)
		return
	})
}

func (r *Requester) setBody(req *http.Request, body Body) {
	if r, ok := body.(io.ReadCloser); !ok {
		req.Body = r
	}

	req.Body = ioutil.NopCloser(body)
}

// Private func that will set the basic auth for a request, will not error
func (r *Requester) setBasicAuth(req *http.Request, basicAuth BasicAuth) {
	req.SetBasicAuth(basicAuth.username, basicAuth.password)
}
