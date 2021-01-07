package ancestor

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	lson := lowestCommonAncestor(root.Left, p, q)
	rson := lowestCommonAncestor(root.Right, p, q)
	if lson != nil && rson != nil {
		return root
	}
	if rson != nil {
		return rson
	}
	return lson
}
