package requester

import "net/http"

// Opt represents an option entry for an http request
type Opt interface{}

// Opts reresents the options entry for an http request
type Opts []Opt

// Modifier is the  modifier func that will modify a request
type Modifier func(*http.Request, *http.Client) error
