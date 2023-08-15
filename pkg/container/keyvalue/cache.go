package keyvalue

import (
	"errors"
	"sync"
	"time"

	"github.com/ponder2000/cache/pkg/core"
)

type mapCache struct {
	sync.RWMutex

	expirationDuration time.Duration
	items              map[string]*core.Item
}

func NewMapCache(expirationDuration time.Duration) *mapCache {
	return &mapCache{items: map[string]*core.Item{}, expirationDuration: expirationDuration}
}

func (m *mapCache) unsafeGetExpiredKeys(expirationTime time.Time) []string {
	expiredKeys := make([]string, 0)
	for k, item := range m.items {
		if item.IsExpired(expirationTime) {
			expiredKeys = append(expiredKeys, k)
		}
	}
	return expiredKeys
}

func (m *mapCache) Purge() int {
	m.RLock()
	expiredKeys := m.unsafeGetExpiredKeys(time.Now())
	m.RUnlock()

	if len(expiredKeys) == 0 {
		return 0
	}

	m.Lock()
	for _, k := range expiredKeys {
		delete(m.items, k)
	}
	m.Unlock()
	return len(expiredKeys)
}

func (m *mapCache) Size() int {
	m.RLock()
	defer m.RUnlock()

	expiredKeys := m.unsafeGetExpiredKeys(time.Now())
	return len(m.items) - len(expiredKeys)
}

func (m *mapCache) SizeAll() int {
	m.RLock()
	defer m.RUnlock()

	return len(m.items)
}

// Add adds the key and value pair if no unexpired item exist else gives error
func (m *mapCache) Add(key string, data any) error {
	m.RLock()
	item, ok := m.items[key]
	m.RUnlock()
	if ok && !item.IsExpired(time.Now()) {
		return errors.New("cache already has the key")
	}

	var e error
	item, e = core.NewItemPtr(data, m.expirationDuration)
	if e != nil {
		return e
	}

	m.Lock()
	m.items[key] = item
	m.Unlock()
	return nil
}

// Set adds the key and value pair if not exist and overwrites if exists
func (m *mapCache) Set(key string, data any) error {
	item, e := core.NewItemPtr(data, m.expirationDuration)
	if e != nil {
		return e
	}

	m.Lock()
	m.items[key] = item
	m.Unlock()
	return nil
}

func (m *mapCache) Get(s string) (any, error) {
	m.RLock()
	item, ok := m.items[s]
	m.RUnlock()
	if !ok {
		return nil, errors.New("key does not exist")
	}

	if item.IsExpired(time.Now()) {
		return nil, errors.New("this data is expired")
	}
	return item.GetData(), nil
}

func (m *mapCache) GetUnexpired() (map[string]any, error) {
	m.RLock()
	defer m.RUnlock()

	unExpired := make(map[string]any)

	now := time.Now()
	for k, item := range m.items {
		if !item.IsExpired(now) {
			unExpired[k] = item.GetData()
		}
	}

	return unExpired, nil
}
