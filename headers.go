package requester

// NewHeaders will return a new instance of HeadersMap
func NewHeaders(entries ...Header) (headers Headers) {
	headers = make(Headers)
	headers.Add(entries...)
	return
}

// Headers is a key/val map representation of the headers of an http request
type Headers map[string]string

// Add will add headers to the headers map
func (h Headers) Add(entries ...Header) {
	for _, header := range entries {
		h[header.Key] = header.Val
	}

	return
}

// ForEach will iterate through ALL entries in an instance of Headers
func (h Headers) ForEach(fn func(key, val string) error) (err error) {
	for key, val := range h {
		if err = fn(key, val); err != nil {
			return
		}
	}

	return
}

// Header is a helper struct for creating a HeadersMap entry
type Header KV
