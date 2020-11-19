# 0002.俩数相加

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

```
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
```

## 解法

1. 使用头插法对每个新建立的节点插入，最后返回即可

### code

```go
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
```
