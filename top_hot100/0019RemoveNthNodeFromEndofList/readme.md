# 19. 删除链表的倒数第N个节点

使用快慢指针

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 首先先定义一个指针向前移动n位
	first := head
	for n != 0 {
		first = first.Next
		n--
	}
	// 如果为nil 则为第一个元素
	if first == nil {
		return head.Next
	}
	// 在定义一个指针second，当first.Next为nil时，则second.Next = second.Next.Next
	second := head
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return head
}
```