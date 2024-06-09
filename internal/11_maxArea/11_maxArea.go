package leetcode_11_maxArea

func maxArea(height []int) int {
	var (
		left, right = 0, len(height) - 1
		res         int
	)
	for left < right {
		res = max(res, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
