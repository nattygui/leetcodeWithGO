# 104. 二叉树的最大深度
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
```
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
```

## 解法

递归方法

go
```go
// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	num1 := dfs(root.Left)
	num2 := dfs(root.Right)
	return max(num1, num2) + 1
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
```

python
```py
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        return self.dfs(root)
    def dfs(self, root):
        if root == None:
            return 0
        num1 = self.dfs(root.left)
        num2 = self.dfs(root.right)
        return max(num1, num2) + 1
```