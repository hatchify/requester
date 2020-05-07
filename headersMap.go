package requester

// NewHeadersMap will return a new instance of HeadersMap
func NewHeadersMap(headerKeyVals ...Header) (headersMap HeadersMap) {
	headersMap = make(HeadersMap)
	headersMap.Add(headerKeyVals...)
	return
}

// HeadersMap is a key/val map representation of the headers of an http request
type HeadersMap map[string]string

// Add will add headers to the headers map
func (h HeadersMap) Add(headerKeyVals ...Header) {
	for _, header := range headerKeyVals {
		h[header.Key] = header.Val
	}

	return
}

// ForEach will iterate through ALL records in an instance of Records
func (h HeadersMap) ForEach(fn func(headerKey, headerVal string) error) (err error) {
	for headerKey, headerVal := range h {
		fn(headerKey, headerVal)
	}

	return
}

// Header is a helper struct for creating a HeadersMap entry
type Header struct {
	Key string
	Val string
}
