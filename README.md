# cache

[![Go Reference](https://pkg.go.dev/badge/github.com/ponder2000/cache.svg)](https://pkg.go.dev/github.com/ponder2000/cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/ponder2000/cache)](https://goreportcard.com/report/github.com/ponder2000/cache)
[![License](https://img.shields.io/github/license/ponder2000/cache)](https://github.com/ponder2000/cache/blob/main/LICENSE)

cache is an in memory cache library to handle caches in the form of key value or FIFO style queue with expiration support.
The main advantage is that we can have one timer to watch multiple cache objects this increases efficiency compared
to `patrickmn/go-cache` when you need to have a large number of cache objects.

## Installation

To install this package, use `go get`:

```bash
go get -u github.com/ponder2000/cache
````

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/ponder2000/cache"
)

func main() {
	// suppose you want to create one janitor for 100 cache where the janitors will clean all 100 caches expired item after 3 seconds
	numberOfCachePerJanitor := 100
	janitorCleanupSchedule := time.Second * 3
	janitorOption := cache.NewJanitorConfigOption(numberOfCachePerJanitor, janitorCleanupSchedule)
	janitorPool := cache.NewJanitorPool(janitorOption)

	// now keep creating cache object as much as you want and register it to the janitor pool
	for i := 0; i<1000; i++ {
		cache := cache.NewListCacher(time.Second * 1)
		janitorPool.RegisterCache(cache)
	}
	
	// note : in the above case janitor pool will have 10 janitor looking after 100 caches each
	
}
```
## [Flow](https://www.figma.com/file/0jA2n6fxuDfxCHiIlUrfv1/Cache-Library?type=whiteboard&node-id=0%3A1&t=nWli2bmx50Zhb5F7-1)

<img src="https://drive.google.com/uc?export=view&id=1CwEtFDuhtP5JdukmXIxdo98WntNlUC_r">

## Documentation

For full package documentation, visit [pkg.go.dev](https://pkg.go.dev/github.com/ponder2000/cache).







## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

Inspired from `patrickmn/go-cache`


## Contact

- Jay Saha
- Twitter: [@chotathanos](https://twitter.com/chotathanos)

Feel free to reach out if you have any questions or feedback!
