package _242ValidAnagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sbyte := []byte(s)
	tbyte := []byte(t)

	smap := make(map[byte]int, 0)
	tmap := make(map[byte]int, 0)

	for i := 0; i < len(sbyte); i++ {
		if _, ok := smap[sbyte[i]]; ok {
			smap[sbyte[i]]++
		} else {
			smap[sbyte[i]] = 0
		}

		if _, ok := smap[tbyte[i]]; ok {
			tmap[tbyte[i]]++
		} else {
			tmap[tbyte[i]] = 0
		}
	}

	for key, value := range smap {
		if tv, ok := tmap[key]; !ok {
			return false
		} else if tv != value {
			return false
		}
	}
	return true
}