package _219ContainsDuplicatesII

func containsNearbyDuplicate(nums []int, k int) bool {
	isExists := make(map[int]int, 0)
	indexmin := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := isExists[nums[i]]; ok {
			if _, ok := indexmin[nums[i]]; ok {
				if indexmin[nums[i]] > i - isExists[nums[i]] {
					indexmin[nums[i]] = i - isExists[nums[i]]
				}
			} else {
				indexmin[nums[i]] = i - isExists[nums[i]]
			}
		}
		isExists[nums[i]] = i

	}

	if len(indexmin) == 0 {
		return false
	}

	for _, v := range indexmin {
		if v > k {
			return false
		}
	}

	return true
}
