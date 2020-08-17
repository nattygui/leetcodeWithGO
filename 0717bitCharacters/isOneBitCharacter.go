package bitcharacters

func isOneBitCharacter(bits []int) bool {

	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == 0 {
			return (len(bits)-i)%2 == 0
		}
	}
	return (len(bits)-1)%2 == 0
}
