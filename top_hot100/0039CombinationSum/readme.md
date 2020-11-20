# 39. 组合总和

给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 

示例 1：

```
输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
```

示例 2：

```
输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
```

## 解法

递归查找，可能存在的值

```go
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	walk(candidates, target, 0, []int{}, &result)
	return result
}

func walk(candidates []int, target int, sum int, curSlice []int, result *[][]int) {
	for i := 0; i < len(candidates); i++ {
		candidate := candidates[i]
		switch {
		case sum+candidate == target:
			oneResult := make([]int, len(curSlice))
			copy(oneResult, curSlice)
			*result = append(*result, append(oneResult, candidate))
			break
		case sum+candidate < target:
			if sum+candidate*2 <= target {
				walk(candidates[i:], target, sum+candidate, append(curSlice, candidate), result)
			} else {
				walk(candidates[i+1:], target, sum+candidate, append(curSlice, candidate), result)
			}
		default:
			break
		}
	}
}
```