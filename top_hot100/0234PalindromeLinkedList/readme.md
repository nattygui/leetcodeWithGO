# 234. 回文链表
请判断一个链表是否为回文链表。

示例 1:
```
输入: 1->2
输出: false
```
示例 2:
```
输入: 1->2->2->1
输出: true
```

## 解法

```go
func isPalindrome(head *ListNode) bool {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    var pre *ListNode
    cur := slow 
    for cur != nil {
        temp := cur.Next
        cur.Next = pre
        pre =cur
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
```