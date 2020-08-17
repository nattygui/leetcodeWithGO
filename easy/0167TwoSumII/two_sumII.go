package _167TwoSumII


// 使用map

func twoSum(numbers []int, target int) []int {
	numsmap := make(map[int]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		numsmap[numbers[i]] = i
	}

	for i := 0; i < len(numbers); i++ {
		temp := target - numbers[i]
		if _, ok := numsmap[temp]; ok {
			if i+1 != numsmap[temp]+1 {
				return []int{i+1, numsmap[temp]+1}
			}
		}
	}

	return nil
}