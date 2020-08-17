package _202HappyNumber

// 使用快慢指针来做
func isHappy1(n int) bool {
	if n <= 0 {
		return false
	}
	fast := next(n)
	slow := n
	for fast != 1 && slow != 1 {
		fast = next(next(fast))
		slow = next(slow)
		if fast == slow {
			return false
		}
	}
	return true
}

// 使用 map
func isHappy2(n int) bool {
	if n <= 0 {
		return false
	}

	search := make(map[int]int, 0)
	search[n] = 0
	temp := n
	for temp != 1 {
		temp = next(temp)
		if _, ok := search[temp]; ok {
			return false
		} else {
			search[temp] = 0
		}
	}
	return true
}

func next(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}