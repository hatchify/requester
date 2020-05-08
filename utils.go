package requester

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Hatch1fy/cookiejar"
)

func getURL(baseURL, path string) (u *url.URL, err error) {
	if u, err = url.Parse(baseURL); err != nil {
		return
	}

	u.Path = path
	return
}

func getCookiesForRequest(jar http.CookieJar) (cookies []*http.Cookie, err error) {
	cookieJar, ok := jar.(*cookiejar.Jar)
	if !ok {
		err = fmt.Errorf("invalid type: expected jar to be of type \"*cookiejar.Jar\", received \"%T\"", jar)
		return
	}

	for _, domainCookies := range cookieJar.Entries {
		cookies = append(cookies, appendCookiesFromDomain(domainCookies)...)
	}

	return
}

func appendCookiesFromDomain(domainCookies map[string]cookiejar.Entry) (cookies []*http.Cookie) {
	for _, entry := range domainCookies {
		cookies = append(cookies, &http.Cookie{Name: entry.Name, Value: entry.Value})
	}

	return
}
