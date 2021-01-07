package reverselist

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

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

func reverseList1(head *ListNode) *ListNode {
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
