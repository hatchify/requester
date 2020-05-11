package requester

import "net/http"

// Modifier is the  modifier func that will modify a request
type Modifier func(*http.Request, *http.Client) error
