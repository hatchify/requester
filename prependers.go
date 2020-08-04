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
		// Create new headers using authorization header
		headers := NewHeaders(authorization)

		// Return options with headers
		return append(o, headers)
	}
}
