# 98. 验证二叉搜索树
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

```
输入:
    2
   / \
  1   3
输出: true
```

示例 2:

```
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
根节点的值为 5 ，但是其右子节点值为 4 。
```

## 解法一

通过前序遍历，获取所有的节点值，然后再判断数组中每个值的前后值

- 时间复杂度: O(n)
- 空间复杂度: O(n)

### code

golang
```go
// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	nums := make([]int, 0)
	dfs(root, &nums)
	if len(nums) <= 1 {
		return true
	}
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
        // 判断条件，若前一个值大于等于当前值，则说明不是二叉搜索树
		if nums[i] <= pre {
			return false
		}
		pre = nums[i]
	}
	return true
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
```py
class Solution:
    def isValidBST(self, root):
        pre_value = float('-inf')
        stack = []
        # 中序遍历非递归方法
        while (root != None) or len(stack) > 0:
            # 获取当前的节点 以及左节点
            while root != None:
                stack.append(root)
                root = root.left
            
            root = stack[-1]
            stack = stack[:-1]
            if root.val <= pre_value:
                return False

            pre_value = root.val
            root = root.right
        
        return True
```

## 解法二

设置每个节点的域值，并且上一个节点做为下一个节点的边界

- 时间复杂度: O(n)
- 空间复杂度: O(n)

### code

golang
```go
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
```

python
```py
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isValidBST(self, root):
        return self.helper(root)
    def helper(self, root, lower=float('-inf'), upper=float('inf')):
        if root == None:
            return True
        if root.val <= lower or root.val >= upper:
            return False
        
        return self.helper(root.left, lower, root.val) and self.helper(root.right, root.val, upper)
```