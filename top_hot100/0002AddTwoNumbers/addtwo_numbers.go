package addtwonumbers

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	result := new(ListNode)                   // 新建头节点
	cur := result                             // 记录当前节点的值，用于替换下一个节点
	for l1 != nil || l2 != nil || carry > 0 { // 判断是否继续
		value := 0
		if l1 != nil {
			value += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			value += l2.Val
			l2 = l2.Next
		}
		value += carry
		carry = value / 10
		cur.Next = &ListNode{ // 新建节点
			Val: value % 10,
		}
		cur = cur.Next // 指向下一个节点
	}
	return result.Next // return 指向的第一个有效节点
}
