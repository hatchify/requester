package mock

// Backend implements storage for requests
type Backend interface {
	Load() (BackendData, error)
	Save(BackendData) error
}
