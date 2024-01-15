func myAtoi(s string) int {
	for i, v := range s {
		if v == ' ' {
			continue
		}
		s = s[i:]
		break
	}

}