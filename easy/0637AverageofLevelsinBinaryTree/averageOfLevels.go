package _637AverageofLevelsinBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先搜索遍历
func averageOfLevelsdfs(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := []float64{}
	levels := make([][]*TreeNode, 0)
	levels = append(levels, []*TreeNode{root})

	i := 0
	for i < len(levels) {
		var curSum float64 = 0.0
		var curNodeSUm float64 = 0.0
		temp := []*TreeNode{}
		for _, node := range levels[i] {

			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
			curSum += float64(node.Val)
			curNodeSUm += 1
		}
		if len(temp) != 0 {
			levels = append(levels, temp)
		}
		res = append(res, curSum/curNodeSUm)
		i++
	}
	return res
}

// 深度度优先搜索遍历
func averageOfLevels(root *TreeNode) []float64 {
	sum := make([]int, 0)
	count := make([]int, 0)
	level := 0
	bfs(root, level, &sum, &count)

	res := make([]float64, len(count))
	for i := 0; i < len(sum); i++ {
		res[i] = float64(sum[i]) / float64(count[i])
	}
	return res
}

func bfs(n *TreeNode, level int, sum *[]int, count *[]int) {
	if n == nil {
		return
	}
	level++
	if len(*sum) < level {
		*sum = append(*sum, 0)
		*count = append(*count, 0)
	}
	(*sum)[level-1] += n.Val
	(*count)[level-1]++
	bfs(n.Left, level, sum, count)
	bfs(n.Right, level, sum, count)
}
