package _599MinimumIndexSumOfTwoLists

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	var res []string
	min := math.MaxInt64
	kv := map[string]int{}
	kv1 := map[string]bool{}
	for index, str := range list1 {
		kv[str] = index
		kv1[str] = false
	}
	for index, str := range list2 {
		if _, ok := kv[str]; ok {
			kv[str] += index
			kv1[str] = true
			if min > kv[str] {
				min = kv[str]
			}
		}
	}

	for k, v := range kv {
		if v == min && kv1[k] == true {
			res = append(res, k)
		}
	}

	return res
}
