package mock

// StoreData is a lookup table for response samples by request samples
type StoreData map[RequestSample]ResponseSample

// NewFlatRecords will return a new flat store from the store data
func (s StoreData) NewFlatRecords() (fs FlatRecords) {
	fs = make(FlatRecords, 0, len(s))
	for request, response := range s {
		fs = append(fs, makeFlatRecord(request, response))
	}

	return
}
