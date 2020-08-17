package _225ImplementStackUsingQueues

type MyStack struct {
	saveValues []int
	index int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		saveValues: make([]int, 0),
		index: 0,
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.saveValues = append(this.saveValues, x)
	this.index++
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.index == 0 {
		return 0
	}
	this.index--
	res := this.saveValues[this.index]
	this.saveValues = this.saveValues[0: this.index]
	return res
}


/** Get the top element. */
func (this *MyStack) Top() int {
	if this.index == 0 {
		return 0
	}
	return this.saveValues[this.index-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.index == 0 {
		return true
	}
	return false
}
