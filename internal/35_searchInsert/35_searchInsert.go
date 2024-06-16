package leetcode_35_searchInsert

func searchInsert(nums []int, target int) int {
	var (
		n     = len(nums)
		left  = 0
		right = n
	)
	for left < right {
		mid := (left + right) >> 1
		val := nums[mid]
		var find bool = val >= target
		if !find {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
