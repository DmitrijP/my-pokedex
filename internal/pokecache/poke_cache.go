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
	mu       sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		entries:  make(map[string]CacheEntry),
		interval: interval,
		mu:       sync.RWMutex{},
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
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

	keysToDelete := make([]string, 0)
	for key, value := range c.entries {
		newTime := value.createdAt.Add(c.interval)
		if newTime.Before(time.Now()) {
			keysToDelete = append(keysToDelete, key)
		}
	}

	for _, key := range keysToDelete {
		delete(c.entries, key)
	}
}
