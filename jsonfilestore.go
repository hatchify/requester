package requester

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// JsonFileStore
type JsonFileStore struct {
	data map[RequestSample]ResponseSample
	file *os.File
	path string
}

type JsonFlatRecord struct {
	Request  RequestSample  `json:"request"`
	Response ResponseSample `json:"response"`
}

type JsonFlatStore []JsonFlatRecord

// NewJsonFileStore creates a new store
func NewJsonFileStore(path string) (s *JsonFileStore){
	var jsonFile *os.File
	var err error

	if jsonFile, err = os.Open(path); err != nil {
		//TODO: implement errors
		fmt.Println("implement errors")
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var flatStore = JsonFlatStore{}
	if err := json.Unmarshal(byteValue, &flatStore); err != nil {
		fmt.Println("couldn't parse the file")
	}

	var data = make(map[RequestSample]ResponseSample)

	for _, v := range flatStore {
		data[v.Request] = v.Response
	}

	s = &JsonFileStore{data, jsonFile, path}
	return
}

func (m *JsonFileStore) GetAll() interface{} {
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

func (m *JsonFileStore) Save() {
	var jsonStore = JsonFlatStore{}

	for request, response := range m.data {
		jsonStore = append(jsonStore,
			JsonFlatRecord{
			Request:  request,
			Response: response,
		})
	}

	byteValue, _ := json.MarshalIndent(jsonStore, "", " ") //TODO: MarshallIndent is just for me
	_ = ioutil.WriteFile(m.path, byteValue, 0644)
}

