package storage

import "errors"

var ErrorNotFound = errors.New("not found")

type DB interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte) error
}
