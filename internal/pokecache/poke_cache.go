package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]CacheEntry
	interval time.Duration
	mu       sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		entries:  make(map[string]CacheEntry),
		interval: interval,
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val, exists := c.entries[key]
	if exists {
		return val.val, exists
	}
	return nil, exists
}

func (c *Cache) ReapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range c.entries {
		newTime := value.createdAt.Add(c.interval)
		if newTime.Before(time.Now()) {
			delete(c.entries, key)
		}
	}
}
