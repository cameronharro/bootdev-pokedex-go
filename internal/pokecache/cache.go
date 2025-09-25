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
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.entries[key]
	if !ok {
		return []byte{}, ok
	}
	return entry.val, ok
}

// func (c *Cache) reapLoop(interval time.Duration) {
// 	ticker := time.NewTicker(interval)
// }

func NewCache(duration time.Duration) *Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
		mutex:   sync.Mutex{},
	}
	// cache.reapLoop(5 * time.Second)
	return &cache
}
