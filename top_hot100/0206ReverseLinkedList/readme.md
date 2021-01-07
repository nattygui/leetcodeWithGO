# 206. 反转链表
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

## 解法一

递归

```go
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	prev, next := newHead, newHead.Next
	for next != nil {
		prev = next
		next = next.Next
	}
	prev.Next = head
	head.Next = nil
	return newHead
}
```

## 解法二

迭代

```go
func reverseList(head *ListNode) *ListNode {
	newHead := (*ListNode)(nil)
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = temp
	}
	return newHead
}
```