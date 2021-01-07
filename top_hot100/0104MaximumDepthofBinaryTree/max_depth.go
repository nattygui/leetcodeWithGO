package maxdepth

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
