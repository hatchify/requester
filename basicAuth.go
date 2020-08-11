package requester

// NewBasicAuth will return a new BasicAuth
func NewBasicAuth(username, password string) (bp BasicAuth) {
	bp.username = username
	bp.password = password
	return
}

// BasicAuth represents the basic auth with username/password
type BasicAuth struct {
	username string
	password string
}
