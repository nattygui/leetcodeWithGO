# 337. 打家劫舍 III
在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

示例 1:
```
输入: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \ 
     3   1

输出: 7 
解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
```

## 解法

```go
// TreeNode 二叉时节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := dfs(root.Left)
	rightMax := dfs(root.Right)
	// 表示选中当前节点
	v1 := root.Val
	if root.Left != nil {
		v1 += dfs(root.Left.Left)
		v1 += dfs(root.Left.Right)
	}
	if root.Right != nil {
		v1 += dfs(root.Right.Left)
		v1 += dfs(root.Right.Right)
	}

	// 表示未选中状态
	v2 := leftMax + rightMax
	return max(v1, v2)
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
```