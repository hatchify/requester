package requester

import (
	"net/url"
)

func getURL(baseURL, path string) (u *url.URL, err error) {
	if u, err = url.Parse(baseURL); err != nil {
		return
	}

	u.Path = path
	return
}
