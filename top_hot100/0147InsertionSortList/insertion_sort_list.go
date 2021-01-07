package sortlist

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail, cur := head, head.Next
	tail.Next = nil
	for cur != nil {
		// 获取当前节点
		temp := cur
		cur = cur.Next // 使当前节点指向下一个

		// 遍历链表节点
		// 如果插入头部
		if temp.Val <= head.Val {
			temp.Next = head
			head = temp
			continue
		}
		// 如果插入尾部
		if temp.Val >= tail.Val {
			tail.Next = temp
			tail = temp
			tail.Next = nil
			continue
		}
		// 插入中间
		prev, next := head, head.Next
		for next != nil {
			if temp.Val >= prev.Val && temp.Val <= next.Val {
				prev.Next = temp
				temp.Next = next
				break
			}
			prev = next
			next = next.Next
		}
	}
	return head
}
