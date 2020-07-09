package mock

func newFlatRecord(req RequestSample, res ResponseSample) (fp *FlatRecord) {
	f := makeFlatRecord(req, res)
	fp = &f
	return
}

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
