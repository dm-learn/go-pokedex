package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	cache map[string]cacheEntry
	interval time.Duration
	mu *sync.Mutex
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	return entry.val, ok
}

// func (c Cache) reapLoop() {

// }

type cacheEntry struct {
	createdAt time.Time
	val []byte
}
