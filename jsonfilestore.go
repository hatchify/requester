package requester

// JsonFileStore
type JsonFileStore struct {
	data map[RequestSample]ResponseSample
}

// NewJsonFileStore creates a new store
func NewJsonFileStore() (s *JsonFileStore){
	s = &JsonFileStore{make(map[RequestSample]ResponseSample)}
	return
}

func (m *JsonFileStore) GetAll() *JsonFileStore {
	return m
}

// Get gets data duuh
func (m *JsonFileStore) Get(request RequestSample) (response ResponseSample, err error) {

	var ok bool

	if response, ok = m.data[request]; !ok {
		err = nil
	}
	return
}

// Set saves data
func (m *JsonFileStore) Set(request RequestSample, response ResponseSample) {
	m.data[request] = response
}

