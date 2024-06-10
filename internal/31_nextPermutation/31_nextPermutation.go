package leetcode_31_nextPermutation

func nextPermutation(nums []int) {
	var (
		n = len(nums)
	)
	flip := func(idx int) {
		cur := nums[idx]
		for i := n - 1; i > idx; i-- {
			if nums[i] > cur {

				nums[i], nums[idx] = nums[idx], nums[i]
				var (
					left, right = idx + 1, n - 1
				)
				for left < right {
					nums[left], nums[right] = nums[right], nums[left]
					left++
					right--
				}
				return
			}
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			continue
		} else {
			flip(i)
			return
		}
	}
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
