package _437PathSumIII

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 递归
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return pathForm(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func pathForm(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	cnt := 0
	if root.Val == sum {
		cnt++
	}

	cnt += pathForm(root.Left, sum-root.Val)
	cnt += pathForm(root.Right, sum-root.Val)

	return cnt
}

// 先序遍历
func pathSumOne(root *TreeNode, sum int) int {
	count := 0
	path := make([]int, 0)
	count = dfs(root, path, sum)
	return count
}

func dfs(root *TreeNode, path []int, sum int) int {
	if root == nil {
		return 0
	}

	count := 0

	for i := 0; i < len(path); i++ {
		path[i] += root.Val
		if path[i] == sum {
			count++
		}
	}
	path = append(path, root.Val)
	if root.Val == sum {
		count++
	}

	leftPath := make([]int, len(path))
	copy(leftPath, path)
	rightPath := make([]int, len(path))
	copy(rightPath, path)

	count += dfs(root.Left, leftPath, sum)
	count += dfs(root.Right, rightPath, sum)

	return count
}

