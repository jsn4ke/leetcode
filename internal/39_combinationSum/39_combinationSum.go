package leetcode_39_combinationSum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	sort.Ints(candidates)
	var dfs func(idx int, remain int, res []int)
	dfs = func(idx int, remain int, res []int) {
		if remain < 0 {
			return
		}
		if remain == 0 {
			ans = append(ans, res)
			return
		}
		for i := idx; i < len(candidates); i++ {
			v := candidates[i]
			if remain < v {
				return
			}
			dfs(i, remain-v, append([]int{v}, res...))
		}
	}
	dfs(0, target, nil)
	return ans
}
