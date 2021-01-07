# 114. 二叉树展开为链表
给定一个二叉树，原地将它展开为一个单链表。

```
例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
```

## 解法

通过中序遍历进行替换

```go
// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	r := root.Right
	if root.Left != nil {
		root.Right = root.Left
		root.Left = nil
	} else {
		root.Right = nil
	}
	flatten(r)
	r1 := root
	for r1.Right != nil {
		r1 = r1.Right
	}
	r1.Right = r
}
```