package greatestletter

func nextGreatestLetter(letters []byte, target byte) byte {
	negMin := byte('z')
	posMin := byte('z')
	posFlag := false
	negFlag := false
	for _, letter := range letters {
		if letter < target && negMin > letter {
			negMin = letter
			negFlag = true
		}

		if letter > target && posMin > letter {
			posMin = letter
			posFlag = true
		}
	}
	if posFlag {
		return posMin
	}

	if negFlag {
		return negMin
	}

	return target
}
