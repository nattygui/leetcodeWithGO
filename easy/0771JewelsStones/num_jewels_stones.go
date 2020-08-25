package numjewelsinstones

func numJewelsInStones(J string, S string) int {
	rubysMap := make(map[byte]bool)
	res := 0

	for i := 0; i < len(J); i++ {
		rubysMap[J[i]] = true
	}

	for i := 0; i < len(S); i++ {
		if rubysMap[S[i]] {
			res++
		}
	}
	return res
}
