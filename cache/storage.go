package cache

type Cache interface {
	Store(id string, content []byte) error
	FetchOne(id string) ([]byte, error)
	KeyExists(id string) (bool, error)
}
