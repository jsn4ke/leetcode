package leatcode

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	if a == e {
		return 1
	}
	if b == f && b != d {
		return 1
	}
	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}
	if abs(c-e) == abs(d-f) {
		if abs(c-a) == abs(d-b) {
			if c < e && a < c && a > e {
				return 1
			}
			if e < c && a < e && a > c {
				return 1
			}
		}
	}
	return 2
}
