package mergetrees

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	dfs(t1, t2)

	return t1
}

func dfs(t1 *TreeNode, t2 *TreeNode) {
	if t1 == nil && t2 != nil {
		t1 = t2
	}
	if t2 == nil {
		return
	}
	if t1.Val < t2.Val {
		t1.Val = t2.Val
	}
	dfs(t1.Left, t2.Left)
	dfs(t1.Right, t2.Right)
}
