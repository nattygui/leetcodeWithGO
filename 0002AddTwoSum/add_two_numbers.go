package _002AddTwoSum

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	out := &ListNode{}
	curNode := out

	next := 0

	for l1 != nil || l2 != nil || next > 0 {
		sum := next
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		next = sum / 10
		curNode.Next = &ListNode{Val: sum % 10}
		curNode = curNode.Next
	}
	return out.Next
}