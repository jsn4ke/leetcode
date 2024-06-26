package leetcode_311_multiply

// a*b b*c - ac
func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	var (
		res = make([][]int, len(mat1))
	)
	for i := 0; i < len(mat1); i++ {
		res[i] = make([]int, len(mat2[0]))

		for j := 0; j < len(mat2[0]); j++ {
			for k := 0; k < len(mat2); k++ {
				res[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}
	return res
}
