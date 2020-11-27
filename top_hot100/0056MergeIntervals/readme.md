# 56. 合并区间

给出一个区间的集合，请合并所有重叠的区间。

示例 1:
```
输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
```

示例 2:
```
输入: intervals = [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间
```

## 解法

1. 首先进行排序，排序的大小按照每个区间的开始进行排序

2. 定义当前的checkIndex ，判断下一个区间与当前的区间是否重叠

3. 若重叠则进行合并， 并从原切片中删除当前的区间，否则checkIndex+1

## code

```go
type sortSliceInt [][]int

func (ss sortSliceInt) Len() int { return len(ss) }
func (ss sortSliceInt) Less(i, j int) bool { return ss[i][0] < ss[j][0] }
func (ss sortSliceInt) Swap(i, j int) { ss[i], ss[j] = ss[j], ss[i] }

func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }
	sort.Sort(sortSliceInt(intervals))
	checkIndex := 0
	for i := 1; i < len(intervals); {
		if intervals[checkIndex][0] <= intervals[i][0] && intervals[checkIndex][1] >= intervals[i][0] {
			intervals[checkIndex][1] = max(intervals[checkIndex][1], intervals[i][1])
			intervals = append(intervals[:i], intervals[i+1:]...)
		} else {
			checkIndex++
			if checkIndex == i {
				i++
			}
		}
	}

	return intervals
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
```

