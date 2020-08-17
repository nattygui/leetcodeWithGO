package _401BinaryWatch

import (
	"fmt"
)

/*
 方法一 回溯法
套模板：
	func tree(选择,路径){
	   结束条件
	   遍历分叉
		 if 满足剪枝条件
			continue
		 进入节点前干啥
		 递归节点
		 遍历节点后干啥
	}
*/

func readBinaryWatch(num int) []string {
	if num == 0 {
		return []string{"0:00"}
	}

	res := make([]string, 0)
	path := make([]int, 0)

	trace(&res, path, num)

	return res
}

func trace(res *[]string, path []int, num int) {
	// 结束条件
	if len(path) == 10 && avaiable(path, num) {
		t := convert(path)
		if t != "" {
			*res = append(*res, t)
		}
		return
	}
	// 遍历分叉
	for _, v := range []int{0, 1} {
		// 满足剪纸条件
		if len(path) >= 10 && !avaiable(path, num) {
			return
		}
		path = append(path, v) // 进入节点前干啥
		trace(res, path, num) // 递归节点
		path = path[:len(path)-1] // 遍历节点后干啥
	}

}

// 判断是否满足有num个1
func avaiable(path []int, num int) bool {
	count := 0
	for _, v := range path {
		if v == 1 {
			count++
		}
	}
	return count == num
}

// 转换为时间字符
func convert(path []int) string {
	hours := 0
	minutes := 0
	timeFormat := "%d:%02d"

	for i := 0; i < 4; i++ {
		if path[i] == 1 {
			hours += 1<<i
		}
	}

	for i := 4; i < 10; i++{
		if path[i] == 1 {
			minutes += 1<<(i-4)
		}
	}

	if hours > 11 || minutes > 59 {
		return ""
	}

	return fmt.Sprintf(timeFormat, hours, minutes)
}


/*
方法二： 暴力解法
*/
func readBinaryWatchOne(num int) []string {
	if num == 0 {
		return []string{"0:00"}
	}
	res := make([]string, 0)
	timeFormat := "%d:%02d"

	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if count(i) + count(j) == num {
				res = append(res, fmt.Sprintf(timeFormat, i, j))
			}
		}
	}
	return res
}

func count(number int) int {
	sum := 0
	for number != 0 {
		if number % 2 == 1 {
			sum++
		}
		number /= 2
	}
	return sum
}