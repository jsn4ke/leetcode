package leatcode

func stringShift(s string, shift [][]int) string {
	var (
		length = len(s)
	)
	for _, v := range shift {
		if 0 == v[0] {
			idx := v[1] % length
		} else {
			idx := length - v[1]%length
		}
		s = s[idx:] + s[:idx]

	}
}
