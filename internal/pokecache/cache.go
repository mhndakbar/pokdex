package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entry: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go cache.reaploop(interval)
	return cache
}

func (c *Cache) reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.entry {
		if v.createdAt.Before(now.Add(-interval)) {
			delete(c.entry, k)
		}
	}
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		value:     value,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	c.mu.Lock()
	entry, ok := c.entry[key]
	defer c.mu.Unlock()
	if !ok {
		return nil, false
	}
	return entry.value, true
}
