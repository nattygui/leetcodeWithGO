package _232ImplementQueueUsingStacks

type MyQueue struct {
	queue []int
	length int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		queue:  []int{},
		length: 0,
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.queue = append(this.queue, x)
	this.length += 1
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.length == 0 {
		return -1
	}

	temp := this.queue[0]
	this.queue = this.queue[1: this.length]
	this.length--

	return temp
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.length == 0 {
		return -1
	}

	return this.queue[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.length == 0 {
		return true
	}
	return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
