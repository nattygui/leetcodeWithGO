package pathsum

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	count := 0
	dfs(root, []int{}, sum, &count)
	return count
}

func dfs(root *TreeNode, nums []int, sum int, count *int) {
	if root == nil {
		return
	}
	nums = append(nums, root.Val)
	newSum := 0
	for j := len(nums) - 1; j >= 0; j-- {
		newSum += nums[j]
		if newSum == sum {
			*count++
		}
	}
	numsLeft := make([]int, len(nums))
	copy(numsLeft, nums)
	numsRight := make([]int, len(nums))
	copy(numsRight, nums)
	dfs(root.Left, numsLeft, sum, count)
	dfs(root.Right, numsRight, sum, count)
}
