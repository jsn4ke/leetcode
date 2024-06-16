package leetcode_33_search

func search(nums []int, target int) int {
	var (
		left  = 0
		n     = len(nums)
		right = n - 1
	)
	for left < right {
		mid := (left + right) >> 1
		val := nums[mid]
		if nums[right] > val {
			right = mid
		} else {
			left = mid + 1
		}
	}
	var (
		l = left
		r = left + n
	)
	for l < r {
		mid := (l + r) >> 1
		val := nums[mid%n]
		if val <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if left == l {
		return -1
	}
	if nums[(l-1)%n] != target {
		return -1
	}
	return (l - 1) % n
}
