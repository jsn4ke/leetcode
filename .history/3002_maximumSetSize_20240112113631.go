package leatcode

import "sort"

func maximumSetSize(nums1 []int, nums2 []int) int {
	var (
		mp    = map[int]int{}
		pairs [][2]int
		size  = len(nums1) + len(nums2)
		half  = size
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
		if pairs[i][1] != pairs[j][1] {
			return pairs[i][1] < pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})
	if len(pairs) >= half {
		return half
	}
	for last := len(pairs) - 1; last >= 0; last-- {
		size -= pairs[last][1]
		if size < half {
			return last + 1
		} else if size == half {
			return last
		}
	}
	return 0
}
