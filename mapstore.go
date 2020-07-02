package requester

// MapStore
type MapStore struct {
	data map[RequestSample]ResponseSample
}

// NewMapStore creates a new store
func NewMapStore() (s *MapStore){
	s = &MapStore{make(map[RequestSample]ResponseSample)}
	return
}

func (m *MapStore) GetAll() *MapStore {
	return m
}

// Get gets data duuh
func (m *MapStore) Get(request RequestSample) (response ResponseSample, err error) {

	var ok bool

	if response, ok = m.data[request]; !ok {
		err = nil
	}
	return
}

// Set saves data
func (m *MapStore) Set(request RequestSample, response ResponseSample) {
	m.data[request] = response
}

