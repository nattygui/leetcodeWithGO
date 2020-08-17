package _190ReverseBits

func reverseBits(num uint32) uint32 {
	var temp1 uint32
	var temp2 uint32

	for i := 0; i < 16; i++ {
		temp1 = (1 << i) & num
		temp2 = (1 << (31 - i)) & num

		num = num &^ (1 << i)
		num = num &^ (1 << (31 - i))

		num = num + temp2 >> (31 - i - i)
		num = num + temp1 << (31 - i - i)
	}
	return num
}
