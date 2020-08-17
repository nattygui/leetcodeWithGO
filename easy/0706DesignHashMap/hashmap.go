package haspmap

//MyHashMap hasmap 结构体
type MyHashMap struct {
	bucket []*Node
}

// Node 存储数据节点
type Node struct {
	key  int
	val  int
	next *Node
}

// put 将当前值放入node节点中
func (node *Node) put(key, value int) *Node {
	if node == nil {
		node = &Node{
			key: key,
			val: value,
		}
		return node
	}
	// 查找是否当前key值
	curNode := node

	// 若含有当前key值 则直接替换
	for curNode != nil {
		if curNode.key == key {
			curNode.val = value
			return node
		}
		curNode = curNode.next
	}

	// 没找到则插入头部
	newNode := &Node{
		key: key,
		val: value,
	}
	newNode.next = node
	node = newNode

	return node
}

// findValue 获取指定key的value值
func (node *Node) findValue(key int) int {
	if node == nil {
		return -1
	}

	// 查找key值 若有则返回value
	curNode := node
	for curNode != nil {
		if curNode.key == key {
			return curNode.val
		}
		curNode = curNode.next
	}
	return -1
}

// remove 根据指定的key 值删除当前value
func (node *Node) remove(key int) *Node {
	if node == nil {
		return node
	}
	// 首先判断第一个node值是否相同
	if node.key == key {
		return node.next
	}
	// 继续判断
	preNode := node
	curNode := node.next
	for curNode != nil {
		if curNode.key == key {
			preNode.next = preNode.next.next
			break
		}
		preNode = curNode
		curNode = curNode.next
	}
	return node
}

// Constructor 创建hashmap
func Constructor() MyHashMap {
	return MyHashMap{
		bucket: make([]*Node, 2069),
	}
}

// Put 放入一个 key，value
func (hashmap *MyHashMap) Put(key int, value int) {
	hash := hashmap.hash(key)

	node := hashmap.bucket[hash].put(key, value)
	hashmap.bucket[hash] = node
}

// Get hashmap 获取指定的key的value
func (hashmap *MyHashMap) Get(key int) int {
	hash := hashmap.hash(key)
	return hashmap.bucket[hash].findValue(key)
}

// Remove 删除指定的key值
func (hashmap *MyHashMap) Remove(key int) {
	hash := hashmap.hash(key)
	node := hashmap.bucket[hash].remove(key)
	hashmap.bucket[hash] = node
}

// hash 获取指定key的hash值
func (hashmap *MyHashMap) hash(key int) int {
	return key % len(hashmap.bucket)
}
