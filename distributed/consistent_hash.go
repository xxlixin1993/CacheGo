package distributed

import (
	"hash/crc32"
	"sort"
	"strconv"
)

func NewContainer(virtualNodes int, fn HashFunc) *Container {
	c := &Container{
		virtual:  virtualNodes,
		hashFunc: fn,
		hashMap:  make(map[int]string),
	}
	if c.hashFunc == nil {
		c.hashFunc = crc32.ChecksumIEEE
	}
	return c
}

// Add nodes to the hashMap
func (c *Container) Add(nodes ...*ContainerNode) {
	for i := 0; i < c.virtual; i++ {
		for _, node := range nodes {
			hash := int(c.hashFunc([]byte(strconv.Itoa(i) + node.nodeName)))
			c.hashMap[hash] = node.nodeName

			// Weighted
			for j := 0; j < node.weight; j++ {
				c.nodes = append(c.nodes, hash)
			}
		}
	}

	sort.Ints(c.nodes)
}

// Find the search cache key in the hashMap
func (c *Container) Get(key string) string {
	if c.IsEmpty() {
		return ""
	}

	hash := int(c.hashFunc([]byte(key)))

	idx := sort.SearchInts(c.nodes, hash)

	// Cycled back
	if idx == len(c.nodes) {
		idx = 0
	}

	return c.hashMap[c.nodes[idx]]
}

// Return the container whether empty or not
func (c *Container) IsEmpty() bool {
	return len(c.nodes) == 0
}
