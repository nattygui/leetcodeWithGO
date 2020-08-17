package _203RemoveLinkedListElements

type ListNode struct {
    Val int
    Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}

	curNode := head.Next
	preNode := head

	for curNode != nil {
		if curNode.Val == val {
			curNode = curNode.Next
			preNode.Next = curNode
		} else {
			curNode = curNode.Next
			preNode = preNode.Next
		}
	}
	return head
}