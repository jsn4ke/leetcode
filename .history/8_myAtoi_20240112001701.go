package leatcode

func myAtoi(s string) int {
	for i, v := range s {
		if v == ' ' {
			continue
		}
		s = s[i:]
		break
	}
	var times = 1
	if s[i] == '+' {
		s = s[1:]
	} else if s[i] == '-' {
		times = -1
		s = s[1:]
	}
	var (
		last int
	)
	for i, v := range s {
		if v > '0' && v < '9' {
			last = last*10 + v - '0'
		} else {
			break
		}
	}
	return times * last
}
