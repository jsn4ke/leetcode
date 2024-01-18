package leetcode_2171_minimumRemoval

import (
	"math"
	"sort"
)

func minimumRemoval(beans []int) int64 {
	var (
		ret   int64 = math.MaxInt64
		total int64
	)
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	sort.Ints(beans)
	for _, v := range beans {
		total += int64(v)
	}
	for i, v := range beans {
		ret = min(ret, total-int64(v)*int64(len(beans)-i))
	}
	return ret
}
