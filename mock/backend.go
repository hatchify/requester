package mock

// Backend implements storage for requests
type Backend interface {
	Load() (StoreData, error)
	Save(StoreData) error
}
