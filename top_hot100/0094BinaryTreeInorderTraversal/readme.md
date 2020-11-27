# 94. 二叉树的中序遍历
给定一个二叉树的根节点 root ，返回它的 中序 遍历。

示例 1：

```
输入：root = [1,null,2,3]
输出：[1,3,2]
```

示例 2：

```
输入：root = []
输出：[]
```

示例 3：

```
输入：root = [1,2]
输出：[2,1]
```

## 解法

golang
```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, result)
	*result = append(*result, root.Val)
	dfs(root.Right, result)
}
```

python
```python
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def inorderTraversal(self, root):
        result = []
        self.dfs(root, result)
        return result
    
    def dfs(self, root, result):
        if root == None:
            return
        self.dfs(root.left, result)
        result.append(root.val)
        self.dfs(root.right, result)
```