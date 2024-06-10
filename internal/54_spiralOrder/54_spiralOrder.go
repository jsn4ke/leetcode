package leetcode_54_spiralOrder

func spiralOrder(matrix [][]int) []int {
	var (
		m, n     = len(matrix), len(matrix[0])
		ans      []int
		rowStart = 0
		rowEnd   = m - 1
		colStart = 0
		colEnd   = n - 1
	)
	for {
		for i := colStart; i <= colEnd; i++ {
			x, y := rowStart, i
			ans = append(ans, matrix[x][y])
		}
		rowStart++
		if rowStart > rowEnd {
			break
		}

		for i := rowStart; i <= rowEnd; i++ {
			x, y := i, colEnd
			ans = append(ans, matrix[x][y])
		}
		colEnd--
		if colStart > colEnd {
			break
		}

		for i := colEnd; i >= colStart; i-- {
			x, y := rowEnd, i
			ans = append(ans, matrix[x][y])
		}
		rowEnd--
		if rowStart > rowEnd {
			break
		}

		for i := rowEnd; i >= rowStart; i-- {
			x, y := i, colStart
			ans = append(ans, matrix[x][y])
		}
		colStart++
		if colStart > colEnd {
			break
		}
	}
	return ans
}
