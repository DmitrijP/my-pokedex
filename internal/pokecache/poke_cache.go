package pokecache

import "time"

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]CacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		entries: make(map[string]CacheEntry),
	}
}

func (c *Cache) Add(key string, val []byte) {
}

func (c *Cache) Get(key string) ([]byte, bool) {
	return nil, false
}

func (c *Cache) reapLoop() {
}
