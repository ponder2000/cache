package cache

import (
	"time"

	"github.com/ponder2000/cache/pkg/container"
	"github.com/ponder2000/cache/pkg/janitor"
)

func NewJanitorConfigOption(cachePerJanitor int, tickerInterval time.Duration) *janitor.ConfigOption {
	return janitor.NewConfigOption(cachePerJanitor, tickerInterval)
}

func NewJanitorPool(option *janitor.ConfigOption) *janitor.Pool {
	return janitor.NewJanitorPool(option)
}

func NewListCacher(expirationDuration time.Duration) container.ListCacher {
	return container.NewListCacher(expirationDuration)
}

func NewMapCacher(expirationDuration time.Duration) container.MapCacher {
	return container.NewMapCacher(expirationDuration)
}
