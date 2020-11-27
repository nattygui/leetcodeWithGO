package inordertraversal

// TreeNode 树的节点
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
