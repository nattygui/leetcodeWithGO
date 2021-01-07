package sortlist

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func merge(head1, head2 *ListNode) *ListNode {
	mergeNode := &ListNode{}
	temp1, temp2, temp := head1, head2, mergeNode.Next
	for temp1 != nil && temp2 != nil {
		if temp1.Val < temp2.Val {
			temp = temp1
			temp1 = temp1.Next
		} else {
			temp = temp2
			temp2 = temp2.Next
		}
		temp.Next = nil
		temp = temp.Next
	}
	if temp1 != nil {
		temp = temp1
	}
	if temp2 != nil {
		temp = temp2
	}
	return mergeNode.Next
}
