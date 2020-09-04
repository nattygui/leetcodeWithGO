package increasing

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	values := make([]int, 0)
	dfs(root, &values)
	res := &TreeNode{}
	construct(res, values)
	return res
}

func dfs(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, values)
	*values = append(*values, root.Val)
	dfs(root.Right, values)
}

func construct(root *TreeNode, values []int) {
	length := len(values)
	for index, value := range values {
		root.Val = value
		if index != length {
			root.Right = &TreeNode{}
			root = root.Right
		}
	}
}
