# 146. LRU 缓存机制
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 

进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？

示例：
```
输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]
```

# 解法

LRU 缓存机制可以通过哈希表辅以双向链表实现，我们用一个哈希表和一个双向链表维护所有在缓存中的键值对。

双向链表按照被使用的顺序存储了这些键值对，靠近头部的键值对是最近使用的，而靠近尾部的键值对是最久未使用的。

哈希表即为普通的哈希映射（HashMap），通过缓存数据的键映射到其在双向链表中的位置。

这样以来，我们首先使用哈希表进行定位，找出缓存项在双向链表中的位置，随后将其移动到双向链表的头部，即可在 O(1) 的时间内完成 get 或者 put 操作。具体的方法如下：

- 对于 get 操作，首先判断 key 是否存在：
    - 如果 key 不存在，则返回 -1−1；

    - 如果 key 存在，则 key 对应的节点是最近被使用的节点。通过哈希表定位到该节点在双向链表中的位置，并将其移动到双向链表的头部，最后返回该节点的值。

- 对于 put 操作，首先判断 key 是否存在：

    - 如果 key 不存在，使用 key 和 value 创建一个新的节点，在双向链表的头部添加该节点，并将 key 和该节点添加进哈希表中。然后判断双向链表的节点数是否超出容量，如果超出容量，则删除双向链表的尾部节点，并删除哈希表中对应的项；

    - 如果 key 存在，则与 get 操作类似，先通过哈希表定位，再将对应的节点的值更新为 value，并将该节点移到双向链表的头部。

上述各项操作中，访问哈希表的时间复杂度为 O(1)O(1)，在双向链表的头部添加节点、在双向链表的尾部删除节点的复杂度也为 O(1)O(1)。而将一个节点移到双向链表的头部，可以分成「删除该节点」和「在双向链表的头部添加节点」两步操作，都可以在 O(1)O(1) 时间内完成。

code
```go
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
```