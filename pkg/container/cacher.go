package container

import (
	"time"

	"github.com/ponder2000/cache/pkg/container/keyvalue"
	"github.com/ponder2000/cache/pkg/container/list"
)

// any struct implementing this interfaces can be used as a cache store

type Cacher interface {
	Purge() int
	Size() int
	SizeAll() int
}

type ListCacher interface {
	Cacher
	Add(any) error
	GetUnexpired() ([]any, error)
}

func NewListCacher(expirationDuration time.Duration) ListCacher {
	return list.NewListCache(expirationDuration)
}

type MapCacher interface {
	Cacher
	Add(string, any) error
	Set(string, any) error
	Get(string) (any, error)
	GetUnexpired() (map[string]any, error)
}

func NewMapCacher(expirationDuration time.Duration) MapCacher {
	return keyvalue.NewMapCache(expirationDuration)
}
