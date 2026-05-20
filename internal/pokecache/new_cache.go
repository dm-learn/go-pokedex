package pokecache

import (
	"time"
	"sync"
)

func NewCache(interval time.Duration) Cache {
	return Cache {
		cache: map[string]cacheEntry{},
		interval: interval,
		mu: &sync.Mutex{},
	}
}
