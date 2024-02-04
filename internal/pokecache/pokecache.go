package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry

	// create mutex to protect the cache from concurrent access
	mux *sync.Mutex
}

type CacheEntry struct {
	value    []byte
	createAt time.Time
}

func NewCache(duration time.Duration) Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(duration)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	// lock the cache to prevent concurrent access
	c.mux.Lock()
	// unlock the cache when the function returns
	defer c.mux.Unlock()

	c.cache[key] = CacheEntry{
		value:    value,
		createAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// lock the cache to prevent concurrent access
	c.mux.Lock()
	// unlock the cache when the function returns
	defer c.mux.Unlock()

	cacheEntry, exists := c.cache[key]
	return cacheEntry.value, exists
}

func (c *Cache) Delete(key string) {
	delete(c.cache, key)
}

func (c *Cache) reap(duration time.Duration) {
	// lock the cache to prevent concurrent access
	c.mux.Lock()
	// unlock the cache when the function returns
	defer c.mux.Unlock()

	cutoffTime := time.Now().UTC().Add(-duration)
	for key, entry := range c.cache {
		// if the entry is older than the cutoff time, delete it
		if entry.createAt.Before(cutoffTime) {
			c.Delete(key)
		}
	}
}

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for range ticker.C {
		c.reap(duration)
	}
}
