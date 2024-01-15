package leetcode_78_subsets

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	var (
		traverse func(idx int, takes ...int)
		ret      [][]int
	)
	traverse = func(idx int, takes ...int) {
		if idx >= len(nums) {
			ret = append(ret, append([]int{}, takes...))
			fmt.Println(takes)
			return
		}
		traverse(idx+1, append(takes, nums[idx])...)
		traverse(idx+1, takes...)
	}
	traverse(0)
	return ret
}
