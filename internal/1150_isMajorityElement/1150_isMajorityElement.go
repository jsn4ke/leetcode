package leetcode_1150_isMajorityElement

func isMajorityElement(nums []int, target int) bool {
	find := func() (int, bool) {
		var (
			l = 0
			r = len(nums)
		)
		for l < r {
			m := l + (r-l)>>1
			if nums[m] < target {
				l = m + 1
			} else {
				r = m
			}
		}
		return l, l != len(nums) && nums[l] == target
	}
	idx, ok := find()
	if !ok {
		return false
	}
	// 4 = 3; 9 = 5
	half := (len(nums))>>1 + 1
	return idx+half <= len(nums) && nums[idx+half-1] == target
}
