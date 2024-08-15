package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.entries[key]
	if ok {
		return entry.val, true
	}

	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		select {
		case <-c.ticker.C:
			c.reapEntries(interval)
		case <-c.done:
			return
		}
	}
}

func (c *Cache) reapEntries(interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	now := time.Now()
	for k, entry := range c.entries {
		if now.Sub(entry.createdAt) > interval {
			delete(c.entries, k)
		}
	}
}
