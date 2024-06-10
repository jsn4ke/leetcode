package leetcode_15_threeSum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var (
		n   = len(nums)
		ans [][]int
	)
	for i := 0; i < n; i++ {
		a := nums[i]
		if a > 0 {
			break
		}
		if 0 != i && a == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			b, c := nums[left], nums[right]
			s := a + b + c
			if s == 0 {
				ans = append(ans, []int{a, b, c})
				left++
				right--
				for left < n && b == nums[left] {
					left++
				}
				for right < n && c == nums[right] {
					right--
				}
			} else if s < 0 {
				left++
			} else if s > 0 {
				right--
			}
		}
	}
	return ans
}
