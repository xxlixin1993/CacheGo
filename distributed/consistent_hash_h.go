package distributed

var hashRing *HashRing

// Hash function can be customized
type HashFunc func(data []byte) uint32

type HashRing struct {
	virtual  int
	nodes    []int
	hashMap  map[int]string
	hashFunc HashFunc
}

type ContainerNode struct {
	nodeName string
	weight   int
}
