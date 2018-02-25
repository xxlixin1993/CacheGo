package lru

import (
	"container/list"
	"github.com/xxlixin1993/CacheGo/logging"
)

// Initialize cache module
func InitCache(maxEntryLen int) {
	LRUCache = NewCache(maxEntryLen)
}

// Create a new Cache struct.
func NewCache(maxEntryLen int) *Cache {
	return &Cache{
		maxEntryLen: maxEntryLen,
		ll:          list.New(),
		cache:       make(map[string]*list.Element),
	}
}

// Add value to the cache.
// Return success or not
func (c *Cache) Add(key string, value string) bool {
	if c.cache == nil {
		logging.Error("[LRU] Add: plz create cache first")
		return false
	}

	// Judge whether it exists or not
	if le, ok := c.cache[key]; ok {
		c.ll.MoveToFront(le)
		le.Value.(*entry).value = value
		return true
	}

	// Add value
	element := c.ll.PushFront(&entry{
		key:   key,
		value: value,
	})
	c.cache[key] = element

	// Judge whether it out of max entry length
	if c.maxEntryLen != 0 && c.ll.Len() > c.maxEntryLen {
		c.Recycling()
	}

	return true
}

// Get the value of the specified index
// Return find value and success or not
func (c *Cache) Get(key string) (value interface{}, success bool) {
	if c.cache == nil {
		logging.Error("[LRU] Get: plz create cache first")
		return nil, false
	}

	// Judge whether it exists or not
	if le, ok := c.cache[key]; ok {
		c.ll.MoveToFront(le)
		return le.Value.(*entry).value, true
	}

	return nil, false
}

// Delete the provided key from the cache.
func (c *Cache) Delete(key string) bool {
	if c.cache == nil {
		logging.Error("[LRU] Delete: plz create cache first")
		return false
	}

	// Judge whether it exists or not
	if le, ok := c.cache[key]; ok {
		pEntry := le.Value.(*entry)
		c.ll.Remove(le)
		delete(c.cache, pEntry.key)
	} else {
		logging.Warning("[LRU] Delete: don not have key: ", key)
	}

	return true
}

// Get length of the cache list.
func (c *Cache) Len() int {
	if c.cache == nil {
		logging.Error("[LRU] Len: plz create cache first")
		return 0
	}

	return c.ll.Len()
}

// Clear the cache.
func (c *Cache) Clear() {
	c.ll = nil
	c.cache = nil
}

// Recycling, remove old cache
func (c *Cache) Recycling() {
	if c.cache == nil {
		logging.Error("[LRU] Recycling: plz create cache first")
		return
	}

	removeLen := c.ll.Len() - c.maxEntryLen
	for i := 0; i < removeLen; i++ {
		c.ll.Back()
	}
}
