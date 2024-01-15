package leatcode

import "sort"

func maximumSetSize(nums1 []int, nums2 []int) int {
	var (
		mp    = map[int]int{}
		pairs [][2]int
	)
	for _, v := range nums1 {
		mp[v]++
	}
	for _, v := range nums2 {
		mp[v]++
	}
	for k, v := range mp {
		pairs = append(pairs, [2]int{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] < pairs[j][1] {
			return true
		}
	})

}
