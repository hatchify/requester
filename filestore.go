package requester

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// FileStore
type FileStore struct {
	data map[RequestSample]ResponseSample
	file *os.File
	path string
}

type FlatRecord struct {
	Request  RequestSample  `json:"request"`
	Response ResponseSample `json:"response"`
}

type FlatStore []FlatRecord

// NewFileStore creates a new store
func NewFileStore(path string) (s *FileStore) {
	var jsonFile *os.File
	var err error

	if jsonFile, err = os.Open(path); err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var flatStore = FlatStore{}
	if err := json.Unmarshal(byteValue, &flatStore); err != nil {
		fmt.Println("couldn't parse the file")
	}

	var data = make(map[RequestSample]ResponseSample)

	for _, v := range flatStore {
		data[v.Request] = v.Response
	}

	s = &FileStore{data, jsonFile, path}
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
	var jsonStore = FlatStore{}

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
