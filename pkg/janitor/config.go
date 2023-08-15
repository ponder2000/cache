package janitor

import "time"

type ConfigOption struct {
	cachePerJanitor int
	tickerInterval  time.Duration
}

func NewConfigOption(cachePerJanitor int, tickerInterval time.Duration) *ConfigOption {
	return &ConfigOption{cachePerJanitor: cachePerJanitor, tickerInterval: tickerInterval}
}
