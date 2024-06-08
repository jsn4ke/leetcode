package leetcode_55_canJump

func canJump2(nums []int) bool {
	var (
		size = len(nums)
	)
	if size == 0 {
		return false
	}
	if size == 1 {
		return true
	}
	nums[size-1] = 1
	for last := size - 2; last >= 0; last-- {
		step := nums[last]
		nums[last] = 0
		for i := 1; i < size && i <= step; i++ {
			if 0 != nums[last+i] {
				nums[last] = 1
				break
			}
		}
	}
	return 1 == nums[0]
}

func canJump(nums []int) bool {
	var k int
	for i, v := range nums {
		if i > k {
			return false
		}
		k = max(k, i+v)
	}
	return true
}
