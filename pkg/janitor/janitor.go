package janitor

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ponder2000/cache/pkg/container"
)

type Janitor struct {
	sync.RWMutex
	*log.Logger

	caches []container.Cacher
}

func newJanitor(janitorIndex int) *Janitor {
	j := &Janitor{
		Logger: log.New(os.Stdout, fmt.Sprintf("[Janitor [%d]] ", janitorIndex), log.Lmsgprefix|log.Ldate|log.Ltime),
		caches: make([]container.Cacher, 0),
	}
	return j
}

func (j *Janitor) run(tickerInterval time.Duration) {
	ticker := time.NewTicker(tickerInterval)

	for {
		<-ticker.C
		j.RLock()
		for _, cache := range j.caches {
			_ = cache.Purge()
			//j.Printf("purged %d items from %p", purgedItem, cache)
		}
		j.RUnlock()

		if len(ticker.C) > 0 {
			j.Println("Unhealthy Janitor option. Ticker rate is more than cache purging rate")
		}
	}
}

func (j *Janitor) registerCache(cache container.Cacher) {
	j.Lock()
	defer j.Unlock()
	j.caches = append(j.caches, cache)
}

func (j *Janitor) hasCapacity(maxCaches int) bool {
	j.RLock()
	defer j.RUnlock()

	if len(j.caches) < maxCaches {
		return true
	}
	return false
}
