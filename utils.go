package requester

import "net/url"

func getURL(baseURL, path string) (u *url.URL, err error) {
	if u, err = url.Parse(baseURL); err != nil {
		return
	}

	u.Path = path
	return
}

// kv is a helper struct for key/val pairs
type kv struct {
	Key string
	Val string
}
