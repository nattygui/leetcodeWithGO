package _141LinkedListCycle

type ListNode struct {
	Val int
	Next *ListNode
}

// 使用 map 用来存储键值对
func hasCycleOne(head *ListNode) bool {
	var mapNode = make(map[*ListNode]int, 0)

	for head != nil {
		if _, exists := mapNode[head]; exists {
			return true
		} else {
			mapNode[head] = 1
			head = head.Next
		}
	}
	return false
}

// 双指针（快慢指针）来实现
func hasCycleTwo(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func hasCycleThird(head *ListNode) bool {
	if head == nil {
		return false
	}
	low := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		low = low.Next
		fast = fast.Next.Next
		if low == fast {
			return true
		}
	}
	return false
}
