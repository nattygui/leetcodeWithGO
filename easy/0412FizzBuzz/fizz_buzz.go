package _412FizzBuzz

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string, 0)

	for i := 1; i <= n; i++ {
		if i % 3 != 0 && i % 5 != 0 {
			res = append(res, strconv.Itoa(i))
		} else if i % 3 == 0 && i % 5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i % 3 == 0 {
			res = append(res, "Fizz")
		} else {
			res = append(res, "Buzz")
		}
	}

	return res
}