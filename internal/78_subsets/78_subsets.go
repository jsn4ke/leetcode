package leetcode_78_subsets

import (
	"fmt"
	"sort"
)

func subsets1(nums []int) [][]int {
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

func subsets(nums []int) [][]int {
	var (
		n   = len(nums)
		ans [][]int
		dfs func(idx int, res ...int)
	)
	dfs = func(idx int, res ...int) {
		ans = append(ans, res)
		for i := idx + 1; i < n; i++ {
			var tmp []int
			tmp = append(tmp, res...)
			tmp = append(tmp, nums[i])
			dfs(i, tmp...)
		}
	}
	ans = append(ans, nil)
	for i := 0; i < n; i++ {
		dfs(i, nums[i])
	}
	return ans
}
