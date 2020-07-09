package mock

// NewMapBackend will return a new MapBackend
func NewMapBackend() (fp *MapBackend) {
	var f MapBackend
	fp = &f
	return
}

// MapBackend is a file-based backend
type MapBackend struct{}

// Load will load
func (m *MapBackend) Load() (s StoreData, err error) {
	s = make(StoreData)
	return
}

// Save will persist the data to disk
func (m *MapBackend) Save(s *Store) (err error) {
	return
}
