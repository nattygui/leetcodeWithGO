package _160IntersectionOfTwoLinkedLists

type ListNode struct {
	Val int
	Next *ListNode
}

// 双指针法
// 是对的总会相遇，不是对的及时走到尽头也不会相遇
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	pB := headB
	for pA != pB {
		if pA == nil && pB == nil {
			return nil
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
