package subarraysum

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := make(map[int]int, 0)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if v, ok := m[pre-k]; ok {
			count += v
		}
		m[pre]++
	}
	return count
}

/* 时间复杂度 O(n^2)
func subarraySum(nums []int, k int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			result++
		}
		if i > 0 {
			nums[i] += nums[i-1]
			if nums[i] == k {
				result++
			}
		}
	}

	for i := 1; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i-1] == k {
				result++
			}
		}
	}

	return result
}
*/

/* 超出内存限制
func subarraySum(nums []int, k int) int {
	result := 0
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
		if dp[i][i] == k {
			result++
		}
		if i > 0 {
			dp[0][i] = nums[i] + dp[0][i-1]
			if dp[0][i] == k {
				result++
			}
		}
	}

	for i := 1; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			dp[i][j] = dp[i][j-1] + nums[i]
			if dp[i][j] == k {
				result++
			}
		}
	}

	return result
}
*/
