package leetcode_56_merge

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})
	var (
		start = intervals[0][0]
		end   = intervals[0][1]
		ret   [][]int
	)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < len(intervals); i++ {
		var (
			s, e = intervals[i][0], intervals[i][1]
		)
		if s <= end {
			end = max(end, e)
		} else {
			ret = append(ret, []int{start, end})
			start, end = s, e
		}
	}
	ret = append(ret, []int{start, end})
	return ret
}
