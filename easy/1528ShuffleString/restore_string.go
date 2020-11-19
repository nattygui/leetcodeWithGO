package restorestring

func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))
	for v, i := range indices {
		res[i] = s[v]
	}
	return string(res)
}
