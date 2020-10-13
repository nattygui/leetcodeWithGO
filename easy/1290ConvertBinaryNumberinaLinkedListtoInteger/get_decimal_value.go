package getdecimalvalue

// ListNode 链表结点
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	newHead := reverseNode(head)
	result := 0
	n := 1
	for newHead != nil {
		result += n * newHead.Val
		n *= 2
		newHead = newHead.Next
	}
	return result
}

func reverseNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}

	return pre
}
