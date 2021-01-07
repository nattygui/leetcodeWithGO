package palindrome

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var pre *ListNode
	cur := fast
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}
