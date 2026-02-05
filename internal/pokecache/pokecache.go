package pokecache

import (
	"time"
	"sync"
	)

type Cache struct { 
    entries  map[string] cacheEntry
    interval time.Duration
    mu       sync.Mutex
}

type cacheEntry struct {
    createdAt time.Time
    val       []byte
}


func NewCache(interval time.Duration) *Cache {
    c := &Cache{
        entries:  make(map[string]cacheEntry),
        interval: interval,
    }

    go c.reapLoop()

    return c
}

func (c *Cache) reapLoop() {
    ticker := time.NewTicker(c.interval)
    defer ticker.Stop()

    for range ticker.C {
        c.mu.Lock()
        for key, entry := range c.entries {
            if time.Since(entry.createdAt) > c.interval {
                delete(c.entries, key)
            }
        }
        c.mu.Unlock()
    }
}


func (c *Cache) Add(key string, val []byte) {
	
    	entry := cacheEntry{
        createdAt: time.Now(),
        val:       val,
    	}
    c.mu.Lock()
    defer c.mu.Unlock()
    c.entries[key] = entry
}
    


func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    entry, ok := c.entries[key]
    if !ok {
	return nil, false
	}
    return entry.val, true
}
