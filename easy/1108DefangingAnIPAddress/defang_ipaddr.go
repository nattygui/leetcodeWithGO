package defangipaddr

func defangIPaddr(address string) string {
	sourceLength := len(address)
	addLength := 6
	resultLength := sourceLength + addLength
	addr := make([]byte, resultLength)
	for i := sourceLength - 1; i >= 0; i-- {
		if address[i] != '.' {
			addr[i+addLength] = address[i]
			continue
		}
		addr[i+addLength] = ']'
		addr[i+addLength-1] = '.'
		addr[i+addLength-2] = '['
		addLength -= 2
	}
	return string(addr)
}
