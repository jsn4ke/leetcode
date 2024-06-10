package leetcode_74_searchMatrix

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	var (
		m, n = len(matrix), len(matrix[0])
	)
	idx := sort.Search(m*n, func(i int) bool {
		x, y := i/n, i%n
		return matrix[x][y] >= target
	})
	if idx == m*n || matrix[idx/n][idx%n] != target {
		return false
	}
	return true
}
