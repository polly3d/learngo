package episode0

import (
	"errors"
)

//ErrorNotFound error
var ErrorNotFound = errors.New("not found")

//HashTable is the interface for a simple hash table.
type HashTable interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}
