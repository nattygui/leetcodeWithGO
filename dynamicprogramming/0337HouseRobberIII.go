package dynamicprogramming

func rob3(root *TreeNode) int {
	return postRob(root)
}

func postRob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := postRob(root.Left)
	rightMax := postRob(root.Right)
	// 抢当前位置
	v1 := root.Val
	if root.Left != nil {
		v1 += postRob(root.Left.Left)
		v1 += postRob(root.Left.Right)
	}
	if root.Right != nil {
		v1 += postRob(root.Right.Right)
		v1 += postRob(root.Right.Left)
	}
	// 不抢当前位置
	v2 := leftMax + rightMax

	return max(v1, v2)
}
