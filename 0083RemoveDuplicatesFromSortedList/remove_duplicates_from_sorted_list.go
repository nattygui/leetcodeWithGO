package _083RemoveDuplicatesFromSortedList

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil && head.Next == nil {
		return head
	}
	prevNode := head
	postNode := head.Next
	for postNode != nil {
		if prevNode.Val == postNode.Val {
			prevNode.Next = postNode.Next
			postNode = postNode.Next
		} else {
			prevNode = postNode
			postNode = postNode.Next
		}
	}
	return head
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	tempNode := head
	tempValue := tempNode.Val

	for tempNode.Next != nil {
		if tempNode.Next.Val == tempValue {
			tempNode.Next = tempNode.Next.Next
		} else {
			tempNode = tempNode.Next
			tempValue = tempNode.Val
		}
	}
	return head
}
