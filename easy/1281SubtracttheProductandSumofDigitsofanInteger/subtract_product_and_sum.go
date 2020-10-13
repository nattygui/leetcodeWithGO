package subtractproductandsum

func subtractProductAndSum(n int) int {
	sum := 0
	div := 1

	for n != 0 {
		sum += n % 10
		div *= n % 10
		n /= 10
	}

	return div-sum
}