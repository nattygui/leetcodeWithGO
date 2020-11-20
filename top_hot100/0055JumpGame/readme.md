# 55. 跳跃游戏

给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

示例 1:

```
输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
```

示例 2:

```
输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
```

## 解法一

通过维护一个可到达最大值，来判断是否可到达最后值

- 时间复杂度： O(n)

- 空间复杂度：O(1)

```go
func canJump(nums []int) bool {
    length := len(nums)
    // 定义一个可到达位置
	maxRight := 0
	for i, num := range nums {
        // 若当前的最大值小于当前的位置，说明是不可达的
		if i > maxRight {
			continue
        }
        // 获取最大的可到达位置
        maxRight := max(maxRight, i+nums[i])
        // 若当前最大值已超过 （数组的长度 - 1），则是可到达的
		if maxRight >= length - 1 {
			return true
		}
    }
    // 直接返回不可到达
	return false
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
```

## 解法二 

若不可到达，则原数组必存在一个0值，且0值之前的数最大可到达的位置为0存在的位置

- 时间复杂度: O(n)~O(n^2)

- 空间复杂度: O(1)

```go

func canJump1(nums []int) bool {
	if len(nums) == 1 {
		return true
    }
    // 首先查找是否存在0值
	lastZeroValueIndex := -1
	for j := len(nums) - 2; j >= 0; j-- {
		if nums[j] == 0 {
			lastZeroValueIndex = j
			break
		}
    }
    // 若不存在，则是可到达的
	if lastZeroValueIndex == -1 {
		return true
	}

    // 若存在零值， 将0值之前的不可到达点，全部变为0值
	for j := lastZeroValueIndex - 1; j >= 0; j-- {
		if nums[j] == 0 {
			continue
		}
		arrived := false
		start := j + 1
		end := nums[j] + j
		for step := end; step >= start; step-- {
			if step >= len(nums)-1 || nums[step] != 0 {
				arrived = true
				break
			}
		}
		if !arrived {
			nums[j] = 0
		}
	}
    // 最后判断当前 从索引0开始 有几个零值，若存在则不可到达
	zeroNum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			break
		}
		zeroNum++
	}
	if zeroNum > 0 {
		return false
	}

	return true
}
```