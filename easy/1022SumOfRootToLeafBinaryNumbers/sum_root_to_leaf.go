package sumroottoleaf

// TreeNode 二叉树的节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(r *TreeNode, s int) int {
	if r == nil {
		return 0
	}
	s = s*2 + r.Val
	if r.Left == nil && r.Right == nil {
		return s
	}
	return dfs(r.Left, s) + dfs(r.Right, s)
}
