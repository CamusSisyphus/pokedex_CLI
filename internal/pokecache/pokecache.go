package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	m     *sync.Mutex
}

func NewCache(interval time.Duration) Cache {

	c := Cache{
		cache: make(map[string]CacheEntry),
		m:     &sync.Mutex{},
	}

	go c.reaploop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.m.Lock()
	defer c.m.Unlock()
	c.cache[key] = NewCacheEntry(val)

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	data, ok := c.cache[key]

	return data.val, ok
}

func (c *Cache) reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)

	}
}

func (c *Cache) reap(interval time.Duration) {

	c.m.Lock()
	defer c.m.Unlock()
	timeBeforeInterval := time.Now().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(timeBeforeInterval) {
			delete(c.cache, k)
		}
	}
}
