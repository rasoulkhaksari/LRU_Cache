package LRU

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	nodes    map[int]*Node
}

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

func NewCache(capacity int) LRUCache {
	var cache LRUCache
	cache.capacity = capacity
	cache.nodes = make(map[int]*Node)
	return cache
}

func (lruc *LRUCache) Remove(node *Node) {
	next_node := node.next
	prev_node := node.prev
	if prev_node != nil && next_node != nil {
		next_node.prev = prev_node
		prev_node.next = next_node
	}
}

func (lruc *LRUCache) Get(key int) int {
	result := -1
	node, ok := lruc.nodes[key]
	if ok {
		result = node.val
		lruc.Remove(node)
		lruc.AddToFront(node)
	}
	return result
}

func (lruc *LRUCache) Put(key int, value int) {
	node, ok := lruc.nodes[key]
	if ok {
		lruc.Remove(node)
		node.val = value
		lruc.AddToFront(node)
	} else {
		if len(lruc.nodes) == lruc.capacity {
			delete(lruc.nodes, lruc.tail.key)
			lruc.Remove(lruc.tail)
		}
		var new_node Node
		new_node.key = key
		new_node.val = value
		lruc.nodes[key] = &new_node
		lruc.AddToFront(&new_node)
	}
}

func (lruc *LRUCache) AddToFront(node *Node) {
	if lruc.head == nil {
		lruc.tail = node
		lruc.head = node
	} else {
		lruc.head.prev = node
		node.next = lruc.head
		lruc.head = node
		node.prev = nil
	}
}
