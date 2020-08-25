package rotateddigits

// DigitsMap 对应数字翻转后的数字
var DigitsMap = map[int]int{
	0: 0,
	1: 1,
	8: 8,
	2: 5,
	5: 2,
	6: 9,
	9: 6,
}

func rotatedDigits(N int) int {
	res := 0
	i := 1
	for i <= N {
		if isPerfectNumber(i) {
			res++
		}
		i++
	}
	return res
}

func isPerfectNumber(n int) bool {
	if n < 10 {
		if _, ok := DigitsMap[n]; !ok {
			return false
		}
		if n == DigitsMap[n] {
			return false
		}
		return true
	}
	nums := []int{}
	for n != 0 {
		nums = append(nums, n%10)
		n /= 10
	}

	rotatedNums := []int{}
	for _, num := range nums {
		if _, ok := DigitsMap[num]; !ok {
			return false
		}
		rotatedNums = append(rotatedNums, DigitsMap[num])
	}

	equalLen := 0
	for i := 0; i < len(nums); i++ {
		if rotatedNums[i] != nums[i] {
			return true
		}
		equalLen++
	}
	if equalLen == len(nums) {
		return false
	}

	return true
}

// 解法2 动态规划
func rotatedDigits2(N int) int {
	ans, d := 0, make([]int, N+1)
	copy(d, []int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1})
	for i := 0; i <= N; i++ {
		if d[i/10] == -1 || d[i%10] == -1 {
			d[i] = -1
		} else if d[i] = d[i/10] | d[i%10]; d[i] == 1 {
			ans++
		}
	}
	return ans
}
