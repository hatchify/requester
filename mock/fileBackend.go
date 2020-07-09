package mock

import (
	"encoding/json"
	"fmt"
	"os"
)

// NewFileBackend will return a new FileBackend
func NewFileBackend(path string) (fp *FileBackend) {
	var f FileBackend
	f.path = path
	fp = &f
	return
}

// FileBackend is a file-based backend
type FileBackend struct {
	path string
}

// Load will load
func (m *FileBackend) Load() (s StoreData, err error) {
	var jsonFileBackend *os.File
	if jsonFileBackend, err = os.Open(m.path); err != nil {
		return
	}

	var FlatRecords FlatRecords
	if err = json.NewDecoder(jsonFileBackend).Decode(&FlatRecords); err != nil {
		return
	}

	s = FlatRecords.NewStoreData()
	return
}

// Save will persist the data to disk
func (m *FileBackend) Save(s StoreData) (err error) {
	var f *os.File
	if f, err = os.OpenFile(m.path, os.O_RDWR, 0744); err != nil {
		err = fmt.Errorf("error opening backend FileBackend: %v", err)
		return
	}
	defer f.Close()

	// Set FileBackend index to 0
	if err = f.Truncate(0); err != nil {
		err = fmt.Errorf("error truncating FileBackend: %v", err)
		return
	}

	// Create flat store from store data
	fs := s.NewFlatRecords()

	// Encode flat store as JSON to disk
	return json.NewEncoder(f).Encode(fs)
}