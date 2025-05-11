package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mutex   sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mutex.Unlock()

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	entry, ok := c.entries[key]
	c.mutex.Unlock()

	if ok {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			<-ticker.C
			c.mutex.Lock()

			now := time.Now()

			for key, entry := range c.entries {
				if now.Sub(entry.createdAt) > interval {
					delete(c.entries, key)
				}
			}
			c.mutex.Unlock()
		}
	}()
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
	}

	cache.reapLoop(interval)
	return cache

}
