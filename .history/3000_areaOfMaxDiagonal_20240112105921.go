package leatcode

func areaOfMaxDiagonal(dimensions [][]int) int {
	var (
		line = 0
		s    = 0
	)
	square := func(a, b int) int {
		return a*a + b*b
	}
	for _, v := range dimensions {
		sq := square(v[0], v[1])
		if line < sq {
			line = sq
		}
	}
}
