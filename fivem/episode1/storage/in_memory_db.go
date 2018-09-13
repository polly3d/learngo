package storage

import (
	"sync"
)

type inMemoryDB struct {
	m map[string][]byte
	l sync.RWMutex
}

//NewInMemoryDB create a db
func NewInMemoryDB() DB {
	return &inMemoryDB{m: make(map[string][]byte)}
}

func (i *inMemoryDB) Get(key string) ([]byte, error) {
	i.l.RLock()
	defer i.l.RUnlock()
	val, ok := i.m[key]
	if !ok {
		return nil, ErrorNotFound
	}
	return val, nil
}

func (i *inMemoryDB) Set(key string, val []byte) error {
	i.l.Lock()
	defer i.l.Unlock()
	i.m[key] = val
	return nil
}
