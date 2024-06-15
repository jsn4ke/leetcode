package leetcode_2779_maximumBeauty

import "slices"

func maximumBeauty(nums []int, k int) (ans int) {
	slices.Sort(nums)
	left := 0
	for right, x := range nums {
		for x-nums[left] > k*2 {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
