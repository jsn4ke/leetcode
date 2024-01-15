package leatcode

func stringShift(s string, shift [][]int) string {
	var (
		length = len(s)
	)
	for _, v := range shift {
		idx := v[1] % length
		if 0 == v[0] {

			s = s[idx:] + s[:idx]
		} else {

		}
	}
}
