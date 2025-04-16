// pokecache/cache.go
package pokecache

import (
	"sync"
	"time"
)

// cacheEntry defines a struct to hold the cached data and creation time
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache defines a struct for the cache system
type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

// NewCache creates a new cache with a given expiration interval
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	// Start the reapLoop in a separate goroutine
	go c.reapLoop(interval)

	return c
}

// Add adds a new entry to the cache
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

// Get retrieves an entry from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

// reapLoop periodically removes expired cache entries
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for key, value := range c.cache {
		if value.createdAt.Before(now.Add(-last)) {
			delete(c.cache, key)
		}
	}
}
