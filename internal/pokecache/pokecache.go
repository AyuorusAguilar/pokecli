package pokecache

import (
	"fmt"
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry
	mu sync.Mutex
}

func (c *Cache) Add(key string, val []byte) error {
	if _, exists := c.cacheMap[key]; exists {
		c.mu.Lock()
		defer c.mu.Unlock()
		return fmt.Errorf("Entry already exists!")
	}

	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, exists := c.cacheMap[key]; exists {
		return entry.value, exists 
	}
	return nil, false
}

func (c *Cache) CleaningLoop(interval time.Duration)  {
	tic := time.NewTicker(interval)
	for range tic.C {
		c.mu.Lock()
		for key, val := range c.cacheMap {
			if time.Since(val.createdAt) > interval {
				delete(c.cacheMap, key)
			} 
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache{
	var c = Cache {
		cacheMap: map[string]cacheEntry{},
		mu: sync.Mutex{},
	}
	go c.CleaningLoop(interval)
	return &c
}