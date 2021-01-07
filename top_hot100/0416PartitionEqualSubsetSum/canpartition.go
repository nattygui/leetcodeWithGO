package partition

// 01背包问题
// 一维矩阵实现
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	searchNum := sum / 2
	bp := make([]bool, searchNum+1)
	bp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := searchNum; j-nums[i] > 0; j-- {
			if bp[searchNum] {
				return true
			}
			bp[j] = bp[j] || bp[j-nums[i]]
		}
	}
	return bp[searchNum]
}

/*
// 0-1背包问题
// 二维矩阵实现
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	searchNum := sum / 2
	// 背包矩阵
	bp := make([][]bool, len(nums))
	// 初始化
	for i := 0; i < len(nums); i++ {
		bp[i] = make([]bool, searchNum+1)
		bp[i][0] = true
	}
	for j := 1; j <= searchNum; j++ {
		if nums[0] == j {
			bp[0][j] = true
		}
	}
	// 遍历每个节点
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= searchNum; j++ {
			if j < nums[i] {
				bp[i][j] = bp[i-1][j]
			} else {
				bp[i][j] = bp[i-1][j] || bp[i-1][j-nums[i]]
			}
		}
	}
	return bp[len(nums)-1][searchNum]
}
*/

/* 暴力解法
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	searchNum := sum / 2
	return hasSubArray(nums, searchNum)
}

func hasSubArray(nums []int, target int) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}
	result := false
	for i, num := range nums {
		result = result && hasSubArray(nums[i+1:], target-num)
	}
	return result
}
*/
