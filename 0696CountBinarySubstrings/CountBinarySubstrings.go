package _696CountBinarySubstrings

type loc struct {
	pre  int
	next int
	num  int
}

func (l *loc) findAll(s string) int {
	s1, s2 := s[l.pre], s[l.next]
	for l.pre > -1 && l.next < len(s) {
		if s1 == s[l.pre] && s2 == s[l.next] {
			l.num++
			l.pre--
			l.next++
		} else {
			return l.num
		}
	}
	return l.num
}

func countBinarySubstrings(s string) int {
	res := 0
	i := 1
	for i < len(s) {
		if s[i] != s[i-1] {
			newLoc := &loc{
				pre:  i - 1,
				next: i,
				num:  0,
			}
			res += newLoc.findAll(s)
		}
		i++
	}
	return res
}
