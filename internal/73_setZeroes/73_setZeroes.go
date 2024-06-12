package leetcode_73_setZeroes

func setZeroes(matrix [][]int) {
	var (
		m   = len(matrix)
		n   = len(matrix[0])
		col = map[int]bool{}
		row = map[int]bool{}
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for k := range row {
		for i := 0; i < n; i++ {
			matrix[k][i] = 0
		}
	}
	for k := range col {
		for i := 0; i < m; i++ {
			matrix[i][k] = 0
		}
	}
}
