package _021MergeTwoSortedLists

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := new(ListNode)
	temp := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Val = l1.Val
			l1 = l1.Next
		} else {
			temp.Val = l2.Val
			l2 = l2.Next
		}
		temp.Next = new(ListNode)
		temp = temp.Next
	}
	if l1 != nil {
		*temp = *l1
	}
	if l2 != nil {
		*temp = *l2
	}
	return result
}