package leetcode_118_generate

func generate(numRows int) [][]int {
	var (
		ans [][]int
	)
	ans = append(ans, []int{1})
	if 1 == numRows {
		return ans
	}
	ans = append(ans, []int{1, 1})
	if 2 == numRows {
		return ans
	}
	for i := 2; i < numRows; i++ {
		var tmp []int
		tmp = append(tmp, 1)
		last := ans[i-1]
		for j := 0; j < len(last)-1; j++ {
			tmp = append(tmp, last[j]+last[j+1])
		}
		tmp = append(tmp, 1)
		ans = append(ans, tmp)
	}
	return ans
}
