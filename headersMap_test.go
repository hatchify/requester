package requester

import (
	"testing"
)

const (
	contentTypeKey   = "Content-Type"
	contentLengthKey = "Content-Length"
	xAPIKeyKey       = "X-API-Key"

	contentTypeVal   = "application/json"
	contentLengthVal = "1337"
	xAPIKeyVal       = "Jenny_8675309"
)

func Test_NewHeadersMap(t *testing.T) {
	headersMap := NewHeadersMap(
		Header{Key: contentTypeKey, Val: contentTypeVal},
		Header{Key: contentLengthKey, Val: contentLengthVal},
		Header{Key: xAPIKeyKey, Val: xAPIKeyVal},
	)

	if l := len(headersMap); l != 3 {
		t.Fatalf("invalid length: expected length of 3, received %d", l)
	}

	contentType, ok := headersMap[contentTypeKey]
	if !ok {
		t.Fatalf("\"%s\" header was not set", contentTypeKey)
	}

	if contentType != contentTypeVal {
		t.Fatalf("invalid contentTypeVal: expected \"%s\", received \"%s\"", contentTypeVal, contentType)
	}

	contentLength, ok := headersMap[contentLengthKey]
	if !ok {
		t.Fatalf("\"%s\" header was not set", contentLengthKey)
	}

	if contentLength != contentLengthVal {
		t.Fatalf("invalid contentLengthVal: expected \"%s\", received \"%s\"", contentLengthVal, contentLength)
	}

	xAPIKey, ok := headersMap[xAPIKeyKey]
	if !ok {
		t.Fatalf("\"%s\" header was not set", xAPIKeyKey)
	}

	if xAPIKey != xAPIKeyVal {
		t.Fatalf("invalid xAPIKeyVal: expected \"%s\", received \"%s\"", xAPIKeyVal, xAPIKey)
	}
}
