package minstack

import "math"

type MinStack struct {
	stack    []int
	minStack []int
	size     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
		size:     0,
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[this.size]
	this.minStack = append(this.minStack, min(top, x))
	this.size++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.size-1]
	this.minStack = this.minStack[:this.size]
	this.size--
}

func (this *MinStack) Top() int {
	return this.stack[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[this.size]
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
