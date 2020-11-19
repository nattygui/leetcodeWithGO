package twosum

func twoSum(nums []int, target int) []int {
	// 1 构建map索引
	numsMap := make(map[int]int)
	for index, num := range nums {
		numsMap[num] = index
	}
	// 2 遍历数组中的值num，并查找map中是否含有对应的（target-num）值
	result := make([]int, 0, 2)
	for index, num := range nums {
		if v, ok := numsMap[target-num]; ok && index != v {
			result = append(result, index, v)
			break
		}
	}
	return result
}
