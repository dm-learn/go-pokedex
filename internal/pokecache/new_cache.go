package pokecache

import (
	"time"
	"sync"
)

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache {
		entries: map[string]cacheEntry{},
		mu: &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}
