package lru

import "container/list"

type Cache struct {
	// Max entry length, 0 means no limit.
	maxEntryLen int

	// LRU list
	ll *list.List

	// Cache
	cache map[string]*list.Element
}

type entry struct {
	key   string
	value string
}
