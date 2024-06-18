package leetcode_1086_highFive

import "sort"

func highFive(items [][]int) [][]int {
	sort.Slice(items, func(i, j int) bool {
		left, right := items[i], items[j]
		if left[0] != right[0] {
			return left[0] < right[0]
		}
		return left[1] >= right[1]
	})
	var ans [][]int
	for i := 0; i < len(items); {
		key := items[i][0]
		var sum int
		for j := i; j < i+5; j++ {
			sum += items[j][1]
		}
		sum /= 5
		ans = append(ans, []int{key, sum})
		for ; i < len(items) && items[i][0] == key; i++ {
		}
	}
	return ans
}
