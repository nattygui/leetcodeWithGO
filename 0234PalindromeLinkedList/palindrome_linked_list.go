package _234PalindromeLinkedList

type ListNode struct {
    Val int
    Next *ListNode
}

// 利用辅助数组来实现
func isPalindromeOne(head *ListNode) bool {
	numbers := make([]int, 0)
	for head != nil {
		numbers = append(numbers, head.Val)
		head = head.Next
	}

	length := len(numbers)

	for i := 0; i < length / 2; i++ {
		if numbers[i] != numbers[length-i-1] {
			return false
		}
	}
	return true
}

// 使用快慢指针
// 1、找到中间指针
// 2、反转后半部分指针
// 3、前半指针 与 后半指针一一比值
func isPalindromeTwo(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var pre *ListNode
	slow := head
	fast := head


	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	for slow != nil {
		temp := slow
		slow = slow.Next
		temp.Next = pre
		pre = temp
	}

	for head != nil && pre != nil {
		if head.Val != pre.Val {
			return false
		}
		head = head.Next
		pre = pre.Next
	}
	return true
}