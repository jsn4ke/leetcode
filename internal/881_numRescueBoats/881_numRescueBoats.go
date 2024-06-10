package leetcode_881_numRescueBoats

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	var (
		left, right = 0, len(people) - 1
		ans         int
	)
	for left <= right {
		if people[right]+people[left] <= limit {
			right--
			left++
		} else {
			right--
		}
		ans++
	}
	return ans
}
