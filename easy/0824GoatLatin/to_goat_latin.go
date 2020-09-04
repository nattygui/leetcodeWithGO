package goatlatin

import "strings"

func toGoatLatin(S string) string {
	words := strings.Split(S, " ")
	for i, word := range words {
		app := []byte{}
		for j := 0; j < i+1; j++ {
			app = append(app, 'a')
		}
		temp := []byte(word)
		switch temp[0] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			temp = append(temp, []byte{'m', 'a'}...)
			temp = append(temp, app...)
		default:
			temp = append(temp[1:], temp[0])
			temp = append(temp, []byte{'m', 'a'}...)
			temp = append(temp, app...)
		}
		words[i] = string(temp)
	}
	return strings.Join(words, " ")
}
