package requester

import (
	"net/url"
)

// NewQuery will return a new instance of Query
func NewQuery(entries ...QueryParam) (query Query) {
	query = make(Query)
	query.Add(entries...)
	return
}

// Query is a key/val map representation of the query of an http request
type Query map[string]string

// Add will add headers to the headers map
func (q Query) Add(entries ...QueryParam) {
	for _, query := range entries {
		q[query.Key] = query.Val
	}

	return
}

// ForEach will iterate through ALL entries in an instance of Query
func (q Query) ForEach(fn func(key, val string) error) (err error) {
	for key, val := range q {
		fn(key, val)
	}

	return
}

// Encode will encode the query
func (q Query) Encode() string {
	var query = url.Values{}
	q.ForEach(func(key, val string) (err error) {
		query.Add(key, val)
		return
	})

	return query.Encode()
}

// QueryParam is a helper struct for creating a Query entry
type QueryParam kv
