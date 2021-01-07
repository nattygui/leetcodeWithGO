package findtargetsumways

// 二维01背包解法
func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if S >= 0 && sum < S {
		return 0
	}
	if S < 0 && -sum < S {
		return 0
	}
	length := sum*2 + 1
	bp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		bp[i] = make([]int, length)
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			bp[i][sum+nums[i]]++
			bp[i][sum-nums[i]]++
			continue
		}
		for j := 0; j < length; j++ {
			if j+nums[i] < length && j+nums[i] >= 0 {
				bp[i][j] += bp[i-1][j+nums[i]]
			}
			if j-nums[i] < length && j-nums[i] >= 0 {
				bp[i][j] += bp[i-1][j-nums[i]]
			}
		}
	}
	return bp[len(nums)-1][S+sum]
}

/* 递归暴力解法
func findTargetSumWays(nums []int, S int) int {
    if len(nums) == 0 {
        return 0
    }
    count := 0
    dfs(nums, 0, 0, S, &count)
    return count
}

func dfs(nums []int, index int, sum int, S int, count *int) {
    if len(nums) == index {
        if sum == S {
            *count++
        }
        return
    }

    dfs(nums, index+1, sum+nums[index], S, count)
    dfs(nums, index+1, sum-nums[index], S, count)
}
*/
