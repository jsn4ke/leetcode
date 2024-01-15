package leatcode

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	if a == e {
		if a != c {
			return 1
		}
		if b < f && (d < b || d > f) {
			return 1
		}
		if f < b && (d < f || d > b) {
			return 1
		}
	}
	if b == f {
		if b != d {
			return 1
		}
		if a < e && (c < a || c > e) {
			return 1
		}
		if e < a && (c < e || c > a) {
			return 1
		}
	}
	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}
	if abs(c-e) == abs(d-f) {
		if abs(a-e) == abs(b-f) && abs(c-a) == abs(d-b) {
			if c < e && (a < c || a > e) {
				return 1
			}
			if e < c && (a < e || a > c) {
				return 1
			}
		} else {
			return 1
		}
	}
	return 2
}
