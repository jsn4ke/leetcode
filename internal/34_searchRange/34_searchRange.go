package leetcode_34_searchRange

import "sort"

func searchRange(nums []int, target int) []int {
	idx := sort.SearchInts(nums, target)
	if idx == len(nums) {
		return []int{-1, -1}
	}
	next := sort.SearchInts(nums, target+1)
	if idx == len(nums) {
		return []int{idx, len(nums) - 1}
	} else {
		return []int{idx, next - 1}
	}
}
