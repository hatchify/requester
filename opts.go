package requester

import (
	"net/http"
	"net/url"
)

// NewOpts will return a new instance of RequestOpts
func NewOpts(query url.Values, headers HeadersMap, modifiers ...Modifier) (op *Opts) {
	var o Opts
	o.query = query
	o.headers = headers
	op = &o
	return
}

// Opts represents optional parameters for an HTTP Request
type Opts struct {
	query     url.Values
	headers   HeadersMap
	modifiers Modifier
}

// Modifier is the  modifier func that will modify a request
type Modifier func(r *http.Request) (err error)
