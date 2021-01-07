# 155. 最小栈
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。
 
示例:
```
输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
```

## 解法

本题的解题点，如何线性的获取栈的最小值。
可以通过建立一个额外的最小栈来存储每次入栈是当前的最小值。

```go
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
```