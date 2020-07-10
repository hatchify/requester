package mock

// BackendData is a lookup table for response samples by request samples
type BackendData map[RequestSample]ResponseSample

// NewFlatRecords will return a new flat store from the backend data
func (s BackendData) NewFlatRecords() (fs FlatRecords) {
	fs = make(FlatRecords, 0, len(s))
	for request, response := range s {
		fs = append(fs, makeFlatRecord(request, response))
	}

	return
}
