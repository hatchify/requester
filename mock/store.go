package mock

import "fmt"

// NewStore creates a new store
func NewStore(be Backend) (sp *Store, err error) {
	var s Store
	s.be = be
	if s.data, err = s.be.Load(); err != nil {
		return
	}

	sp = &s
	return
}

// Store manages a set of mock requests
type Store struct {
	be   Backend
	data BackendData
}

// Get gets data duuh
func (s *Store) Get(request RequestSample) (response ResponseSample, err error) {
	var ok bool
	if response, ok = s.data[request]; !ok {
		err = fmt.Errorf("request does not exist in Store")
	}

	return
}

// Set saves data
func (s *Store) Set(request RequestSample, response ResponseSample) {
	s.data[request] = response
}

// Save will save
func (s *Store) Save() (err error) {
	return s.be.Save(s.data)
}
