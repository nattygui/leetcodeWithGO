package _299BullsAndCows

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	bulls := 0
	cows := 0

	numsA := make([]int, 10)
	numsB := make([]int, 10)

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			numsA[secret[i] - '0']++
			numsB[guess[i] - '0']++
		}
	}

	for i := 0; i < 10; i++ {
		if numsA[i] > 0 && numsB[i] > 0 {
			cows += min(numsA[i], numsB[i])
		}
	}

	return fmt.Sprintf("%vA%vB", bulls, cows)
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}