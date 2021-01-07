package rob

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
