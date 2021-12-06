package cache

import (
	"sync"
)

// Item represent an item in cache
type Item struct {
	Content []byte
	Expire  int64
}

// a cache is represented by a map
type Cache struct {
	items map[string]Item
	mux   *sync.RWMutex
	Ttl   int64
}
