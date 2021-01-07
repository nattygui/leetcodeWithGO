# 160. 相交链表
编写一个程序，找到两个单链表相交的起始节点。

## 解法

```go
// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	pB := headB
	for {
		if pA == nil && pB == nil {
			break
		}
		if pA == pB {
			break
		}
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}
		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}
	return pA
}
```