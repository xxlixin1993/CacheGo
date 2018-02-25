package distributed

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"

	"github.com/xxlixin1993/CacheGo/configure"
)

// Initialize consistent hash container
func InitConsistentHash(fn HashFunc) error {
	virtualNodes := configure.DefaultInt("hash.virtualNode", 0)
	hashContainer = NewHashContainer(virtualNodes, fn)

	nodeNum := configure.DefaultInt("node.number", 0)
	for i := 0; i < nodeNum; i++ {
		nodeName := configure.DefaultString("hash.node."+string(i)+"host", "")
		weight := configure.DefaultInt("hash.node."+string(i)+"weight", 0)
		if nodeName == "" || weight == 0 {
			return errors.New("node host or weight can not be empty in node config")
		}

		hashContainer.Add(&ContainerNode{
			nodeName: nodeName,
			weight:   weight,
		})
	}
	return nil
}

func NewHashContainer(virtualNodes int, fn HashFunc) *HashContainer {
	c := &HashContainer{
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
func (c *HashContainer) Add(nodes ...*ContainerNode) {
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
func (c *HashContainer) Get(key string) string {
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
func (c *HashContainer) IsEmpty() bool {
	return len(c.nodes) == 0
}

// TODO Resolve adding or removing nodes while running
