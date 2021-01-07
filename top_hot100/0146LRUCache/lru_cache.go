package lru

// DlinkNode 双向链表节点
type DlinkNode struct {
	key   int
	value int
	prev  *DlinkNode
	next  *DlinkNode
}

// LRUCache 缓存器
type LRUCache struct {
	size  int
	cap   int
	cache map[int]*DlinkNode
	head  *DlinkNode
	tail  *DlinkNode
}

func newDlinkNode(key, value int) *DlinkNode {
	return &DlinkNode{
		key:   key,
		value: value,
	}
}

// Constructor 新建一个LRUCache
func Constructor(capacity int) LRUCache {
	newLRU := LRUCache{
		size:  0,
		cap:   capacity,
		cache: map[int]*DlinkNode{},
		head:  newDlinkNode(1, 1),
		tail:  newDlinkNode(1, 1),
	}
	newLRU.head.next = newLRU.tail
	newLRU.tail.prev = newLRU.head
	return newLRU
}

// Get 获取一个值
func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	node := lru.cache[key]
	removeNode(node)
	lru.moveHead(node)
	return node.value
}

// Put 放入一个节点
func (lru *LRUCache) Put(key int, value int) {
	// 首先查找是否存在
	if node, ok := lru.cache[key]; ok {
		lru.cache[key].value = value
		removeNode(node)
		lru.moveHead(node)
		return
	}
	// 不存在则需要插入
	// 检查容量到达阈值
	if lru.size == lru.cap {
		// 删除第一个元素
		node := lru.tail.prev
		delete(lru.cache, node.key)
		removeNode(node)
		lru.size--
	}
	node := newDlinkNode(key, value)
	lru.moveHead(node)
	lru.cache[key] = node
	lru.size++
}

func (lru *LRUCache) moveTail(node *DlinkNode) {
	lru.tail.prev.next = node
	node.next = lru.tail
	node.prev = lru.tail.prev
	lru.tail.prev = node
}

func (lru *LRUCache) moveHead(node *DlinkNode) {
	lru.head.next.prev = node
	node.next = lru.head.next
	node.prev = lru.head
	lru.head.next = node
}

func removeNode(node *DlinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
