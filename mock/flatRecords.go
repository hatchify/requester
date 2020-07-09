package mock

// FlatRecords is a list of FlatRecord values
type FlatRecords []FlatRecord

// NewStoreData will return new store data from the flat store
func (f FlatRecords) NewStoreData() (sd StoreData) {
	sd = make(StoreData, len(f))
	for _, v := range f {
		sd[v.Request] = v.Response
	}

	return
}
