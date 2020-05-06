package requester

import (
	"net/url"

	common "github.com/Hatch1fy/integrations-common"
)

// NewRequestOpts will return a new instance of RequestOpts
func NewRequestOpts(query url.Values, headers common.HeadersMap) (rp *RequestOpts) {
	var r RequestOpts
	r.query = query
	r.headers = headers
	rp = &r
	return
}

// RequestOpts represents optional parameters for an HTTP Request
type RequestOpts struct {
	query   url.Values
	headers common.HeadersMap
}
