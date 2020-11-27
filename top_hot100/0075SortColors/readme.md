# 75. 颜色分类

给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
 

进阶：

你可以不使用代码库中的排序函数来解决这道题吗？
你能想出一个仅使用常数空间的一趟扫描算法吗？
 

示例 1：

```
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
```

示例 2：

```
输入：nums = [2,0,1]
输出：[0,1,2]
```

示例 3：

```
输入：nums = [0]
输出：[0]
```

示例 4：

```
输入：nums = [1]
输出：[1]
```

# 解法一

使用单指针来分别查找0，1值, 但是需要遍历俩次

```go
func sortColors(nums []int) {
	count := 0
	for i, num := range nums {
		if num == 0 {
			nums[i], nums[count] = nums[count], nums[i]
			count++
		}
	}
	for i, num := range nums {
		if num == 1 {
			nums[i], nums[count] = nums[count], nums[i]
			count++
		}
	}
}
```

# 解法二

```go
func sortColors(nums []int) {
	count0, count1 := 0, 1
	for i, num := range nums {
		if num == 0 {
			nums[i], nums[count0] = nums[count0], nums[i]
			if count0 < count1 {
				nums[i], nums[count1] = nums[count1], nums[i]
			}
			count0++
			count1++
		} else if num == 1 {
			nums[i], nums[count1] = nums[count1], nums[i]
			count1++
		}
	}
}
```