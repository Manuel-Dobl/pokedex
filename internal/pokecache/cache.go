package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]CacheEntry
	mu       sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {

	c.mu.Lock()
	defer c.mu.Unlock()

	oneCacheEntry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.cacheMap[key] = oneCacheEntry

}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheMap[key]

	if ok {
		return entry.val, ok
	} else {
		return nil, false
	}

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.cacheMap {
			if time.Since(entry.createdAt) > interval {
				delete(c.cacheMap, key)
			}

		}
		c.mu.Unlock()
	}

}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheMap: make(map[string]CacheEntry),
	}
	go c.reapLoop(interval)
	return c
}
