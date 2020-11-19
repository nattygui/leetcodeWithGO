package canformarray

func canFormArray(arr []int, pieces [][]int) bool {
	hashArr := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		hashArr[arr[i]] = i + 1
	}

	for i := 0; i < len(pieces); i++ {
		if len(pieces[i]) == 1 && hashArr[pieces[i][0]] < 1 {

			return false
		}
		if len(pieces[i]) > 1 {
			for k := 0; k < len(pieces[i])-1; k++ {
				if hashArr[pieces[i][k+1]]-hashArr[pieces[i][k]] != 1 {
					return false
				}
			}
		}
	}
	return true
}
