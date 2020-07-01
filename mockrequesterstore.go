package requester

import "net/http"

// RequesterStore implements storage for requests
type RequesterStore interface {
	Get(request interface{}) (response *http.Response, err error)
	Set(request interface{}, response *http.Response)
}

// MapStore
type MapStore struct {
	data map[interface{}]*http.Response
}

// NewMapStore creates a new store
func NewMapStore() (s *MapStore){
	return
}

// Get gets data duuh
func (m *MapStore) Get(request interface{}) (response *http.Response, err error) {

	var ok bool

	if response, ok = m.data[request]; !ok {
		err = nil
	}
	return
}

// Set saves data
func (m *MapStore) Set(request interface{}, response *http.Response) {
	m.data[request] = response
}
