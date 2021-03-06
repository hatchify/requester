package mock

func makeFlatRecord(req RequestSample, res ResponseSample) (f FlatRecord) {
	f.Request = req
	f.Response = res
	return
}

// FlatRecord is an entry containing Request/Response pairs
type FlatRecord struct {
	Request  RequestSample  `json:"request"`
	Response ResponseSample `json:"response"`
}
