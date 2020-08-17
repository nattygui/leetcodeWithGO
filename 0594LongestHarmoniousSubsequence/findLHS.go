package _594LongestHarmoniousSubsequence

func findLHS(nums []int) int {
	res := 0
	kv := make(map[int]int)


	for _, v := range nums {
		if _, ok := kv[v]; ok {
			kv[v]++
		} else {
			kv[v] = 1
		}
	}

	for k, v := range kv {
		if v1, ok := kv[k-1]; ok {
			if v + v1 > res {
				res = v + v1
			}
		}
		if v1, ok := kv[k+1]; ok {
			if v + v1 > res {
				res = v + v1
			}
		}
	}

	return res
}
