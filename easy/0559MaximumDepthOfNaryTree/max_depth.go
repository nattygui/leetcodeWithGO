package _559MaximumDepthOfNaryTree

type Node struct {
    Val int
    Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	res := 0
	for _, node := range root.Children {
		res = max(res, maxDepth(node))
	}
	return res + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

