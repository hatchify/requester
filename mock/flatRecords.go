package mock

// FlatRecords is a list of FlatRecord values
type FlatRecords []FlatRecord

// NewBackendData will return new store data from the flat store
func (f FlatRecords) NewBackendData() (b BackendData) {
	b = make(BackendData, len(f))
	for _, v := range f {
		b[v.Request] = v.Response
	}

	return
}
