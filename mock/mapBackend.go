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
func (m *MapBackend) Load() (s StoreData, err error) {
	s = make(StoreData)
	return
}

// Save will not persist data
func (m *MapBackend) Save(s StoreData) (err error) {
	return
}
