package levelorder

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	curLevelNode := make([]*TreeNode, 0)
	curLevelNode = append(curLevelNode, root)
	for {
		if len(curLevelNode) == 0 {
			break
		}
		nextLevelNodes := make([]*TreeNode, 0)
		curLevelValues := make([]int, 0, len(curLevelNode))
		for _, oneNode := range curLevelNode {
			curLevelValues = append(curLevelValues, oneNode.Val)
			if oneNode.Left != nil {
				nextLevelNodes = append(nextLevelNodes, oneNode.Left)
			}
			if oneNode.Right != nil {
				nextLevelNodes = append(nextLevelNodes, oneNode.Right)
			}
		}
		curLevelNode = nextLevelNodes
		result = append(result, curLevelValues)
	}
	return result
}
