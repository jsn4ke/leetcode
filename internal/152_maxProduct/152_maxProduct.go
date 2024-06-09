package leetcode_152_maxProduct

func maxProduct(nums []int) int {
	v0 := float64(nums[0])
	maxF, minF, ans := v0, v0, v0
	for _, iv := range nums[1:] {
		v := float64(iv)
		a, b := maxF, minF
		maxF = max(v, a*v, b*v)
		minF = min(v, a*v, b*v)
		ans = max(ans, maxF)
	}
	return int(ans)
}
