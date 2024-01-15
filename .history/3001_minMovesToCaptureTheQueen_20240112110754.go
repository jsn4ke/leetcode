package leatcode

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	if a == e && a != c {
		return 1
	}
	if b == f && b != d {
		return 1
	}
}
