package leafsimilar

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	r1 := findLeafnodeValue(root1)
	r2 := findLeafnodeValue(root2)
	return isSameValue(r1, r2)
}

func findLeafnodeValue(t *TreeNode) []int {
	res := make([]int, 0)

	dfs(t, &res)
	return res
}

func dfs(t *TreeNode, node *[]int) {
	if t == nil {
		return
	}
	dfs(t.Left, node)
	if t.Left == nil && t.Right == nil {
		*node = append(*node, t.Val)
	}
	dfs(t.Right, node)
}

func isSameValue(r1, r2 []int) bool {
	if len(r1) != len(r2) {
		return false
	}
	lens := len(r1)
	for i := 0; i < lens; i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}
