package requester

import "fmt"

// NewAuthBearerHook will return a new auth bearer hook func
func NewAuthBearerHook(apiKey string) func() Opts {
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
