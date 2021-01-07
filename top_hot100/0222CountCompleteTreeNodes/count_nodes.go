package countnodes

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	result := 0
	for {
		if len(queue) == 0 {
			break
		}
		tempqueue := make([]*TreeNode, 0)
		for _, node := range queue {
			result++
			if node.Left != nil {
				tempqueue = append(tempqueue, node.Left)
			}
			if node.Right != nil {
				tempqueue = append(tempqueue, node.Right)
			}
		}
		queue = tempqueue
	}
	return result
}
