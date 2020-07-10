package mock

// NewMapBackend will return a new MapBackend
func NewMapBackend() (fp *MapBackend) {
	var f MapBackend
	fp = &f
	return
}

// MapBackend is a map struct backend
type MapBackend struct{}

// Load will load
func (m *MapBackend) Load() (s BackendData, err error) {
	s = make(BackendData)
	return
}

// Save will not persist data
func (m *MapBackend) Save(s BackendData) (err error) {
	return
}
