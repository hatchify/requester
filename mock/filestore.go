package mock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func NewFileStore(path string) (f *FileStore) {
	f = &FileStore{
		data: make(map[RequestSample]ResponseSample),
		path: path,
	}

	return f
}

// FileStore
type FileStore struct {
	data map[RequestSample]ResponseSample
	path string
}

type FlatRecord struct {
	Request  RequestSample  `json:"request"`
	Response ResponseSample `json:"response"`
}

type FlatStore []FlatRecord

// NewFileStore creates a new store
func (m *FileStore) Load() (err error) {
	var jsonFile *os.File
	if jsonFile, err = os.Open(m.path); err != nil {
		return
	}

	var flatStore FlatStore
	if err = json.NewDecoder(jsonFile).Decode(&flatStore); err != nil {
		return
	}

	for _, v := range flatStore {
		m.data[v.Request] = v.Response
	}

	return
}

func (m *FileStore) GetAll() interface{} {
	return m
}

// Get gets data duuh
func (m *FileStore) Get(request RequestSample) (response ResponseSample, err error) {

	var ok bool

	if response, ok = m.data[request]; !ok {
		err = fmt.Errorf("request does not exist in FileStore")
	}
	return
}

// Set saves data
func (m *FileStore) Set(request RequestSample, response ResponseSample) {
	m.data[request] = response
}

func (m *FileStore) Save() {
	jsonStore := make(FlatStore, 0, len(m.data))
	for request, response := range m.data {
		jsonStore = append(jsonStore,
			FlatRecord{
				Request:  request,
				Response: response,
			})
	}

	byteValue, _ := json.MarshalIndent(jsonStore, "", " ") //TODO: MarshallIndent is just for me
	_ = ioutil.WriteFile(m.path, byteValue, 0644)
}
