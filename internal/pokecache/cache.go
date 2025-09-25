package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mutex   sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
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
	entry, ok := c.entries[key]
	if !ok {
		return []byte{}, ok
	}
	return entry.val, ok
}

func (c *Cache) reap(ch <-chan time.Time, interval time.Duration) {
	for timestamp := range ch {
		for key, entry := range c.entries {
			if timestamp.Sub(entry.createdAt).Milliseconds() > interval.Milliseconds() {
				c.mutex.Lock()
				delete(c.entries, key)
				c.mutex.Unlock()
			}
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go c.reap(ticker.C, interval)
}

func NewCache(duration time.Duration) *Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
		mutex:   sync.Mutex{},
	}
	cache.reapLoop(5 * time.Second)
	return &cache
}
