package _705DesignHashSet

type MyHashSet struct {
	// 当前大小
	size int
	// 容量
	cap int
	// 负载因子
	growFactor   float64
	shrinkFactor float64
	buf          []*Node
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyHashSet {
	return MyHashSet{
		size:         0,
		growFactor:   0.75,
		shrinkFactor: 0.25,
		cap:          16,
		buf:          make([]*Node, 16),
	}
}

func (this *MyHashSet) hashCode(key int) int {
	return key % this.cap
}

func (this *MyHashSet) needGrow() bool {
	return this.size >= int(float64(this.cap)*this.growFactor)
}

func (this *MyHashSet) needShrink() bool {
	return this.size <= int(float64(this.cap)*this.shrinkFactor) && this.cap > 16
}

func (this *MyHashSet) grow() {
	this.cap *= 2
	this.rehash()
}

func (this *MyHashSet) shrink() {
	this.cap /= 2
	this.rehash()
}

func (this *MyHashSet) rehash() {
	newBuf := make([]*Node, this.cap)
	for i := 0; i < len(this.buf); i++ {
		head := this.buf[i]
		for head != nil {
			hash := this.hashCode(head.Val)
			node := &Node{Val: head.Val}
			if newBuf[hash] == nil {
				newBuf[hash] = node
			} else {
				node.Next = newBuf[hash]
				newBuf[hash] = node
			}
			head = head.Next
		}
	}
	this.buf = newBuf
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	this.size++
	if this.needGrow() {
		this.grow()
	}

	hash := this.hashCode(key)
	node := &Node{Val: key}
	if this.buf[hash] == nil {
		this.buf[hash] = node
	} else {
		node.Next = this.buf[hash]
		this.buf[hash] = node
	}
}

func (this *MyHashSet) Remove(key int) {
	hash := this.hashCode(key)
	if this.buf[hash] == nil {
		return
	} else {
		dummy := &Node{Val: 1<<31 - 1}
		dummy.Next = this.buf[hash]
		head := dummy
		for head != nil && head.Next != nil {
			if head.Next.Val == key {
				t := head.Next
				head.Next = head.Next.Next
				t.Next = nil
				this.size--
				break
			}
			head = head.Next
		}
		this.buf[hash] = dummy.Next
	}
	if this.needShrink() {
		this.shrink()
	}
}

func (this *MyHashSet) Contains(key int) bool {
	hash := this.hashCode(key)
	if this.buf[hash] == nil {
		return false
	} else {
		head := this.buf[hash]
		for head != nil {
			if head.Val == key {
				return true
			}
			head = head.Next
		}
	}
	return false
}
