package buddystrings

func buddyStrings(A string, B string) bool {

	flag := false
	if len(A) > 26 {
		flag = true
	} else {
		letterMap := make(map[byte]bool)
		for i := 0; i < len(A); i++ {
			letterMap[A[i]] = true
		}
		if len(letterMap) != len(A) {
			flag = true
		}
	}

	diffIndex := findIndex(A, B)

	if diffIndex == nil {
		return false
	}

	if len(diffIndex) == 0 && flag {
		return true
	}

	if len(diffIndex) != 2 {
		return false
	}

	if A[diffIndex[0]] != B[diffIndex[1]] || A[diffIndex[1]] != B[diffIndex[0]] {
		return false
	}

	return true
}

func findIndex(A string, B string) []int {
	if len(A) != len(B) {
		return nil
	}
	res := []int{}
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			res = append(res, i)
		}
	}
	return res
}
