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
    c.mu.Lock()
    defer c.mu.Unlock()
    for range ticker.C { 
     for i, _ := range c.entries {
       if c.entries[i].createdAt > c.interval{
        delete(c.entries, i)
       }
      } 
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