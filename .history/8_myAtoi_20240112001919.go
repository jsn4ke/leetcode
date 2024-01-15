package leatcode

import "math"

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
		last int64
	)
	for _, v := range s {
		if v > '0' && v < '9' {
			last = last*10 + int64(v-'0')
			if 1 == times {
				if last > math.MaxInt32 {
					return math.MaxInt32
				}
			} else {
				if -last < math.MinInt32 {
					return math.MinInt32
				}
			}
		} else {
			break
		}
	}
	return times * int(last)
}
