package list

import (
	"sync"
	"time"

	"github.com/ponder2000/cache/pkg/core"
)

type listCache struct {
	sync.RWMutex

	expirationDuration time.Duration
	items              []*core.Item
}

func NewListCache(expirationDuration time.Duration) *listCache {
	return &listCache{items: []*core.Item{}, expirationDuration: expirationDuration}
}

func (l *listCache) unsafeGetLatestExpirationIndex() int {
	expirationTime := time.Now()
	expirationIndex := -1
	for index, item := range l.items {
		if item.IsExpired(expirationTime) {
			expirationIndex = index
		} else {
			break
		}
	}
	return expirationIndex
}

func (l *listCache) Purge() int {
	l.RLock()
	expirationIndex := l.unsafeGetLatestExpirationIndex()
	l.RUnlock()
	if expirationIndex < 0 {
		return 0
	}

	l.Lock()
	l.items = l.items[expirationIndex+1:]
	l.Unlock()
	return expirationIndex + 1
}

func (l *listCache) Size() int {
	l.RLock()
	defer l.RUnlock()

	expiredCnt := l.unsafeGetLatestExpirationIndex() + 1
	return len(l.items) - expiredCnt
}

func (l *listCache) SizeAll() int {
	l.RLock()
	defer l.RUnlock()

	return len(l.items)
}

func (l *listCache) Add(d any) error {
	if item, e := core.NewItemPtr(d, l.expirationDuration); e != nil {
		return e
	} else {
		l.Lock()
		defer l.Unlock()

		l.items = append(l.items, item)
		return nil
	}
}

func (l *listCache) GetUnexpired() ([]any, error) {
	l.RLock()
	defer l.RUnlock()

	latestExpiredIndex := l.unsafeGetLatestExpirationIndex()

	res := make([]any, 0)
	for _, item := range l.items[latestExpiredIndex+1:] {
		res = append(res, item.GetData())
	}

	return res, nil
}
