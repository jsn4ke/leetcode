package leetcode_240_searchMatrix

func searchMatrix(matrix [][]int, target int) bool {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
