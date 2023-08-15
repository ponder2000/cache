package janitor

import (
	"log"
	"os"
	"sync"

	"github.com/ponder2000/cache/pkg/container"
)

type Pool struct {
	sync.RWMutex
	*log.Logger

	option   *ConfigOption
	janitors []*Janitor
}

func NewJanitorPool(option *ConfigOption) *Pool {
	return &Pool{
		Logger:   log.New(os.Stdout, "[Janitor pool] ", log.Lmsgprefix|log.Ldate|log.Ltime),
		option:   option,
		janitors: make([]*Janitor, 0),
	}
}

func (jp *Pool) RegisterCache(cache container.Cacher) {
	lastJanitor := jp.getCurrentJanitor()
	if lastJanitor == nil || !lastJanitor.hasCapacity(jp.option.cachePerJanitor) {

		lastJanitor = newJanitor(len(jp.janitors))
		jp.Lock()
		jp.janitors = append(jp.janitors, lastJanitor)
		jp.Unlock()

		jp.Println("Created a new Janitor, total janitors =", len(jp.janitors))

		lastJanitor.registerCache(cache)
		go lastJanitor.run(jp.option.tickerInterval)
	} else {
		lastJanitor.registerCache(cache)
	}
}

func (jp *Pool) getCurrentJanitor() *Janitor {
	jp.RLock()
	defer jp.RUnlock()
	if len(jp.janitors) == 0 {
		return nil
	}
	lastJanitor := jp.janitors[len(jp.janitors)-1]
	return lastJanitor
}
