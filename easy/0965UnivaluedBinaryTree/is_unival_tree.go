package univaltree

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	valueMap := make(map[int]bool)
	valueMap[root.Val] = true

	nodes := []*TreeNode{root.Left, root.Right}
	for len(nodes) != 0 {
		tempNodes := make([]*TreeNode, 0)
		for _, node := range nodes {
			if node == nil {
				continue
			}
			if !valueMap[node.Val] {
				return false
			}
			tempNodes = append(tempNodes, node.Left)
			tempNodes = append(tempNodes, node.Right)
		}
		nodes = tempNodes
	}
	return true
}
