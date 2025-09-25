package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mutex   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) reap(timestamp time.Time, interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, entry := range c.entries {
		if timestamp.After(entry.createdAt.Add(interval)) {
			delete(c.entries, key)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for timestamp := range ticker.C {
		c.reap(timestamp, interval)
	}
}

func NewCache(reapInterval time.Duration) *Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
		mutex:   &sync.Mutex{},
	}
	go cache.reapLoop(reapInterval)
	return &cache
}
