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
}