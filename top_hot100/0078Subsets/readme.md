# 78. 子集

给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

```
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
```

## 解法

通过递归获取到每个num 将其追加即可

golang
```go
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
    emptySet := make([]int, 0)
    result = append(result, emptySet)
	getAll(nums, emptySet, &result)
	return result
}

func getAll(nums []int, oneSet []int, result *[][]int) {
	if len(nums) == 0 {
		return
	}
	for i, num := range nums {
        tempSet := make([]int, len(oneSet)+1)
        copy(tempSet, oneSet)
		tempSet[len(oneSet)] = num
        *result = append(*result, tempSet)
		getAll(nums[i+1:], tempSet, result)
	}
}
```

python
```python
class Solution:
    def subsets(self, nums):
        result = [[]]
        self.getall(nums, [], result)
        return result
    
    def getall(self, nums, one_set, result):
        if len(nums) == 0:
            return 
        for index, num in enumerate(nums):
            temp = one_set[:] # copy一份新的数组
            temp.append(num)
            result.append(temp)
            self.getall(nums[index+1:], temp, result)
```