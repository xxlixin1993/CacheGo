package distributed

const KVirtualNodes = 40

// Hash function can be customized
type HashFunc func(data []byte) uint32

type Container struct {
	virtual  int
	nodes    []int
	hashMap  map[int]string
	hashFunc HashFunc
}

type ContainerNode struct {
	nodeName string
	weight   int
}
