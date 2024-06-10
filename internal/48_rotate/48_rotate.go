package leetcode_48_rotate

func rotate(matrix [][]int) {
	var (
		n = len(matrix)
	)
	// a[row][col+step], a[row+step][col+diff], a[row+diff][col-step], a[row-step][col] =
	// a[row+step][col+diff], a[row+diff][col-step], a[row-step][col], a[row][col+step]
	for idx := 0; idx < n-1; idx++ {
		rowStart := idx
		rowEnd := n - 1 - idx
		colStart := idx
		colEnd := n - 1 - idx

		for step := 0; step < rowEnd-rowStart; step++ {
			a := matrix[rowStart][colStart+step]
			b := matrix[rowStart+step][colEnd]
			c := matrix[rowEnd][colEnd-step]
			d := matrix[rowEnd-step][colStart]

			matrix[rowStart][colStart+step] = d
			matrix[rowStart+step][colEnd] = a
			matrix[rowEnd][colEnd-step] = b
			matrix[rowEnd-step][colStart] = c
		}
	}
}
