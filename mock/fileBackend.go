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
func (m *FileBackend) Load() (b BackendData, err error) {
	var jsonFileBackend *os.File
	if jsonFileBackend, err = os.Open(m.path); err != nil {
		err = fmt.Errorf("error opening json FileBackend: %v", err)
		return
	}

	var stat os.FileInfo
	if stat, err = jsonFileBackend.Stat(); err != nil {
		err = fmt.Errorf("error can't stat json FileBackend: %v", err)
	}

	var flatRecords FlatRecords

	//Avoids json decode errors on new files
	if stat.Size() > 0 {
		if err = json.NewDecoder(jsonFileBackend).Decode(&flatRecords); err != nil {
			err = fmt.Errorf("error decoding json FileBackend: %v", err)
			return
		}
	}

	b = flatRecords.NewBackendData()
	return
}

// Save will persist the data to disk
func (m *FileBackend) Save(b BackendData) (err error) {
	var f *os.File
	if f, err = os.OpenFile(m.path, os.O_RDWR, 0744); err != nil {
		err = fmt.Errorf("error opening json for saving FileBackend: %v", err)
		return
	}
	defer f.Close()

	// Set FileBackend index to 0
	if err = f.Truncate(0); err != nil {
		err = fmt.Errorf("error truncating FileBackend: %v", err)
		return
	}

	// Create flat store from store data
	fs := b.NewFlatRecords()

	// Encode flat store as JSON to disk
	return json.NewEncoder(f).Encode(fs)
}
