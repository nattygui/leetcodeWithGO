# 148. 排序链表

给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

示例 1：

```
输入：head = [4,2,1,3]
输出：[1,2,3,4]
```

## 解法

对链表自顶向下归并排序的过程如下。

- 找到链表的中点，以中点为分界，将链表拆分成两个子链表。寻找链表的中点可以使用快慢指针的做法，快指针每次移动 2 步，慢指针每次移动 1 步，当快指针到达链表末尾时，慢指针指向的链表节点即为链表的中点。

- 对两个子链表分别排序。

- 将两个排序后的子链表合并，得到完整的排序后的链表。可以使用「21. 合并两个有序链表」的做法，将两个有序的子链表进行合并。

上述过程可以通过递归实现。递归的终止条件是链表的节点个数小于或等于 1，即当链表为空或者链表只包含 1 个节点时，不需要对链表进行拆分和排序。

## code

golang
```go
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
```
