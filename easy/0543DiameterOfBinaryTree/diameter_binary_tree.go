package _543DiameterOfBinaryTree

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var max int

func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	recursionAllNode(root)
	return max

}

func recursionAllNode(root *TreeNode) {
	if root == nil {
		return
	}
	recursionAllNode(root.Left)
	recursionSingleNode(root)
	recursionAllNode(root.Right)
}

func recursionSingleNode(root *TreeNode) {
	if root == nil {
		return
	}
	left := recursion(root.Left, 0)
	right := recursion(root.Right, 0)
	if left + right > max {
		max = left + right
	}
}


func recursion(root *TreeNode, l int) int {
	if root == nil {
		return l + 0
	}
	left := recursion(root.Left, l+1)
	right := recursion(root.Right, l+1)

	if left > right {
		return left
	}
	return right
}

// 2

func diameterOfBinaryTreeOne(root *TreeNode) int {
	ans := 1
	dfs(root, &ans)
	return ans - 1

}

func dfs(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, ans)
	right := dfs(root.Right, ans)
	*ans = maxValue(left+right+1, *ans)
	return maxValue(left, right) + 1
}

func maxValue(x, y int) int {
	if x > y {
		return x
	}
	return y
}