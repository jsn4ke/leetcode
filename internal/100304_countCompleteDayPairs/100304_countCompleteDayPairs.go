package leetcode_100304_countCompleteDayPairs

import "sort"

func countCompleteDayPairs(hours []int) int {
	sort.Ints(hours)
	var (
		n   = len(hours)
		ans int
	)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (hours[i]+hours[j])%24 == 0 {
				ans++
			}
		}
	}
	return ans
}
