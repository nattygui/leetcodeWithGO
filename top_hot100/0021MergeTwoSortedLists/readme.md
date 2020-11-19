# 0021. 合并两个有序链表

```go
// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil { // 如果俩个链表都为空，则直接返回
		return nil
	}
	if l1 == nil { // 若链表1为空，则返回链表2
		return l2
	}
	if l2 == nil { // 若链表2为空，则返回链表1
		return l1
	}
	// 按照头插法合并
	result := new(ListNode)
	cur := result      // 当前节点
	var temp *ListNode // 临时节点
	// 关键
	// 当l1 和 l2 都不为空的时候，判断俩个节点值，若小则赋予temp 并将temp追加到cur.Next上
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp = l1
			l1 = l1.Next
		} else {
			temp = l2
			l2 = l2.Next
		}
		cur.Next = temp
		cur = cur.Next
	}
	// 若经过上述后，l1 与 l2 有一个不为空，则直接追加cur.Next
	// 若l1 不为空，则直接追加
	if l1 != nil {
		cur.Next = l1
	}
	// 若l2 不为空，则直接追加
	if l2 != nil {
		cur.Next = l2
	}
	// 返回结果
	return result.Next
}
```