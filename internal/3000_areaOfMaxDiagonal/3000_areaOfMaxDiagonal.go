package leetcode_3000_areaOfMaxDiagonal

func areaOfMaxDiagonal(dimensions [][]int) int {
	var (
		line = 0
		s    = 0
	)
	f1 := func(a, b int) int { return a*a + b*b }
	f2 := func(a, b int) int { return a * b }
	for _, v := range dimensions {
		t1 := f1(v[0], v[1])
		if line > t1 {
			continue
		}
		t2 := f2(v[0], v[1])
		if line == t1 && s > t2 {
			continue
		}
		line = t1
		s = t2
	}
	return s
}
