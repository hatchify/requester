package requester

import "fmt"

// NewAuthBearerPrepender will return a new auth bearer prepender func
func NewAuthBearerPrepender(apiKey string) func() Opts {
	// Set authorization header
	authorization := Header{
		Key: "Authorization",
		Val: fmt.Sprintf("Bearer %s", apiKey),
	}

	return func() (o Opts) {
		// Return options with authorization headers
		return append(o, authorization)
	}
}

// NewXAPIKEYPrepender will return a new X-API-KEY prepender func
func NewXAPIKEYPrepender(apiKey string) func() Opts {
	// Set authorization header
	authorization := Header{
		Key: "X-API-KEY",
		Val: apiKey,
	}

	return func() (o Opts) {
		// Return options with authorization headers
		return append(o, authorization)
	}
}
