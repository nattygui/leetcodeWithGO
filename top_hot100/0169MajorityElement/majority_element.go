package majorityelement

func majorityElement(nums []int) int {
	length := len(nums)
	result := 0
	numMap := make(map[int]int, 0)
	for _, num := range nums {
		if v, ok := numMap[num]; ok {
			numMap[num] = v + 1
			if v+1 >= length/2 {
				result = num
			}
		}
	}
	return result
}
