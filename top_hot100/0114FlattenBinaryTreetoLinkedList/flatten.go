package flatten

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	r := root.Right
	if root.Left != nil {
		root.Right = root.Left
		root.Left = nil
	} else {
		root.Right = nil
	}
	flatten(r)
	r1 := root
	for r1.Right != nil {
		r1 = r1.Right
	}
	r1.Right = r
}
