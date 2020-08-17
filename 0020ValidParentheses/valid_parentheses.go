package _020ValidParentheses

type bracket struct {
	front byte
	end byte
}

var Brackets = []bracket{
	{
		front: '(',
		end: ')',
	},
	{
		front: '[',
		end: ']',
	},
	{
		front: '{',
		end: '}',
	},
}

func isValid(s string) bool {
	if len(s) == 0 && (len(s) % 2 == 1) {
		return false
	}
	var queue = make([]byte, 0)
	for index, _ := range s {
		queue = append(queue, s[index])
		if len(queue) >= 2 {
			front := queue[len(queue)-2]
			end := queue[len(queue)-1]
			temp := bracket{
				front: front,
				end:   end,
			}

			for _, b := range Brackets {
				if temp == b {
					queue = queue[0: len(queue)-2]
					break
				}
			}
		}
	}
	if len(queue) == 0 {
		return true
	}
	return false
}
