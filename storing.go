package storing

// Storing interface
type Storing interface {
	Upload(name string, contentType string, content []byte) (string, error)
	Download(path string) ([]byte, error)
	Provider() (name string)
	Delete(key string) error
}
