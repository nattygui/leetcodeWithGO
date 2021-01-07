package convert

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	dfs(root, &sum)
	return root
}

func dfs(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	dfs(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	dfs(root.Left, sum)
}
