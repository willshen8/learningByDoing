package cache

import (
	"sync"
	"time"
)

// creates a new cache with ttl as the argument
func NewCache(ttl int64) *Cache {
	return &Cache{
		items: make(map[string]Item),
		mux:   &sync.RWMutex{},
		Ttl:   ttl,
	}
}

// check if a particular cache item is expired or not based on ttl set on the cache
func (c *Cache) isExpired(item *Item) bool {
	if item.Expire == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expire
}

// retrieve an item from cache
func (c *Cache) Get(key string) []byte {
	c.mux.RLock()
	defer c.mux.RUnlock()

	item := c.items[key]
	if c.isExpired(&item) { // cache miss
		delete(c.items, key)
		return nil
	}
	return item.Content
}

func (c *Cache) Add(key string, content []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.items[key] = Item{
		Content: content,
		Expire:  time.Now().Add(time.Second * time.Duration(c.Ttl)).UnixNano(),
	}
}
