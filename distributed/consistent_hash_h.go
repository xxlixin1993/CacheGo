package distributed

var hashContainer *HashContainer

// Hash function can be customized
type HashFunc func(data []byte) uint32

type HashContainer struct {
	virtual  int
	nodes    []int
	hashMap  map[int]string
	hashFunc HashFunc
}

type ContainerNode struct {
	nodeName string
	weight   int
}
