package _605CanPlaceFlowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append(flowerbed, 0)
	flowerbed = append([]int{0}, flowerbed...)

	for i := 0 ; i < len(flowerbed)-1; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] != 1 && flowerbed[i+1] != 1 {
			flowerbed[i] = 1
			n--
		}
	}

	if n > 0 {
		return false
	}
	return true
}
