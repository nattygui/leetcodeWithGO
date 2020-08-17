package _206ReverseLinkedList

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var res *ListNode

	val := 0
	for head != nil {
		val = head.Val
		head = head.Next
		temp := new(ListNode)
		temp.Val = val

		// 头插法
		temp.Next = res
		res = temp
	}
	return res
}
