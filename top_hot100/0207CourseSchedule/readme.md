# 207. 课程表
你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]

给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

示例 1:
```
输入: 2, [[1,0]] 
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
```
示例 2:
```
输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
```

## 解法一 广度优先遍历

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 初始化, y轴表示要修的课程，每个y对应的x 都是y的前置课程
	allCourses := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		allCourses[i] = make([]int, 0)
	}
	// 获取每个课程的前置, 并构建入度表
	inDegree := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		i, j := prerequisite[0], prerequisite[1]
		inDegree[i]++
		allCourses[j] = append(allCourses[j], i)
	}
	queue := make([]int, 0)
	for index, in := range inDegree {
		if in == 0 {
			queue = append(queue, index)
		}
	}
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]
		numCourses--
		for _, course := range allCourses[v] {
			inDegree[course]--
			if inDegree[course] == 0 {
				queue = append(queue, course)
			}
		}
	}
	return numCourses == 0
}
```

## 解法二 深度优先遍历

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 初始化, y轴表示要修的课程，每个y对应的x 都是y的前置课程
	allCourses := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		allCourses[i] = make([]int, 0)
	}
	// 获取每个课程的前置
	for _, prerequisite := range prerequisites {
		i, j := prerequisite[0], prerequisite[1]
		allCourses[j] = append(allCourses[j], i)
	}
	revise := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if !(dfs(allCourses, revise, i)) {
			return false
		}
	}

	return true
}

func dfs(allCourses [][]int, revise []int, i int) bool {
	if revise[i] == 1 {
		return false
	}
	if revise[i] == -1 {
		return true
	}
	revise[i] = 1
	for _, v := range allCourses[i] {
		if !(dfs(allCourses, revise, v)) {
			return false
		}
	}
	revise[i] = -1
	return true
}
```
