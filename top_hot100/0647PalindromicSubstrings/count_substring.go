package count

func countSubstrings(s string) int {
	res := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for l := 0; l < len(s)-1; l++ {
		for i := 0; i+l < len(s); i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true
				res++
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = true
					res++
				}
			} else {
				if dp[i+1][j-1] && s[i] == s[j] {
					dp[i][j] = true
					res++
				}
			}
		}
	}
	return res
}
