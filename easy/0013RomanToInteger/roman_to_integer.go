package _013RomanToInteger

type signal struct {
	value int
	level int
}

var Roman = map[byte]signal {
	'M': {
		value: 1000,
		level: 7,
	},
	'D': {
		value: 500,
		level: 6,
	},
	'C': {
		value: 100,
		level: 5,
	},
	'L': {
		value: 50,
		level: 4,
	},
	'X': {
		value: 10,
		level: 3,
	},
	'V': {
		value: 5,
		level: 2,
	},
	'I': {
		value: 1,
		level: 1,
	},
}

func romanToInt(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if i + 1 < len(s) && Roman[s[i]].level < Roman[s[i+1]].level {
			temp := Roman[s[i+1]].value - Roman[s[i]].value
			result = result + temp
			i++
		} else {
			temp := Roman[s[i]].value
			result = result + temp
		}
	}
	return result
}