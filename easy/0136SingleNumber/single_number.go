package _136SingleNumber


func singleNumber(nums []int) int {
	num := 0
	for _, v := range nums {
		num = num ^ v
	}
	return num
}
