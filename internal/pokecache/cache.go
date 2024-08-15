package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	ticker  *time.Ticker
	done    chan bool
	mutex   sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]CacheEntry),
		ticker:  time.NewTicker(interval),
		done:    make(chan bool),
	}

	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Close() {
	c.done <- true
	c.ticker.Stop()
}
