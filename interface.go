package requester

import "net/http"

// Interface needs to implement all needed Requester methods
type Interface interface {
	Request(method, path string, body []byte, opts Opts) (resp *http.Response, err error)

	Get(path string, opts ...Opt) (resp *http.Response, err error)
	Post(path string, body []byte, opts ...Opt) (resp *http.Response, err error)
	Put(path string, body []byte, opts ...Opt) (resp *http.Response, err error)
	Patch(path string, body []byte, opts ...Opt) (resp *http.Response, err error)
	Delete(path string, opts ...Opt) (resp *http.Response, err error)
}
