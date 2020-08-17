#1. Two Sum

Given an array of integers, return __indices__ of the two numbers such that they add up to a specific target.

You may assume that each input would have __exactly__ one solution, and you may not use the same element twice.

__Example:__

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## 解题思路

初始化一个map，里面装入数组numbers的值 与 索引

```go
source := make(map[int]int, 0)
for i, v := range numbers {
    source[v] =i
}
```

然后遍历numbers，通过判断当前map中是否存在对应的值

```go
for i, v := range numbers {
    // 存在,则返回true给ok
    if _, ok := source[target-v]; ok && source[target-v] != i {

        out = append(out, i, source[target-v])
        return out
    }
}
```