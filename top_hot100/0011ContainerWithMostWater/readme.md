# 0011. 盛最多水的容器

给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。

示例 1：

```
输入：[1,8,6,2,5,4,8,3,7]
输出：49 
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
```

示例 2：

```
输入：height = [1,1]
输出：1
```

示例 3：

```
输入：height = [4,3,2,1,4]
输出：16
```

示例 4：

```
输入：height = [1,2,1]
输出：2
```

## 解法一 暴力解法

- 时间复杂度: O(n^2)

- 空间复杂度: O(1)

### code 

```go
func maxArea(height []int) int {
	length := len(height)
	if length < 2 {
		return 0
	}
	result := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			tempArea := min(height[i], height[j]) * (j - i)
			if tempArea > result {
				result = tempArea
			}
		}
	}
	return result
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
```

## 解法二 

- 时间复杂度: O(n)

- 空间复杂度: O(1)

### code

```go
func maxArea(height []int) int {
	i, j, result := 0, len(height)-1, 0
	for i < j {
		if height[i] < height[j] {
			result = max(result, height[i]*(j-i))
			i++
		} else {
			result = max(result, height[j]*(j-i))
			j--
		}
	}
	return result
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
```