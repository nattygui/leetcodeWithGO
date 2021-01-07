package canfinish

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
