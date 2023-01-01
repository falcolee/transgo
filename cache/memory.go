package cache

import (
	"errors"
	"sync"
)

type InMemoryStorage struct {
	content map[string][]byte
	mux     sync.Mutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		content: map[string][]byte{},
		mux:     sync.Mutex{},
	}
}

func (i *InMemoryStorage) Store(id string, content []byte) error {
	i.mux.Lock()
	defer i.mux.Unlock()
	i.content[id] = content
	return nil
}

func (i *InMemoryStorage) FetchOne(id string) ([]byte, error) {
	i.mux.Lock()
	defer i.mux.Unlock()
	result, ok := i.content[id]
	if !ok {
		return nil, errors.New("index not found")
	}
	return result, nil
}

func (i *InMemoryStorage) KeyExists(id string) (bool, error) {
	_, ok := i.content[id]
	return ok, nil
}

var _ Cache = (*InMemoryStorage)(nil)
