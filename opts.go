package requester

import (
	"net/url"

	common "github.com/Hatch1fy/integrations-common"
)

// NewOpts will return a new instance of RequestOpts
func NewOpts(query url.Values, headers common.HeadersMap) (op *Opts) {
	var o Opts
	o.query = query
	o.headers = headers
	op = &o
	return
}

// Opts represents optional parameters for an HTTP Request
type Opts struct {
	query   url.Values
	headers common.HeadersMap
}
