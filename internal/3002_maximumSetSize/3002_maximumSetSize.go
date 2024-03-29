package leetcode_3002_maximumSetSize

func maximumSetSize(nums1 []int, nums2 []int) int {
	var (
		s1   = map[int]struct{}{}
		s2   = map[int]struct{}{}
		half = len(nums1) / 2
	)
	for _, v := range nums1 {
		s1[v] = struct{}{}
	}
	for _, v := range nums2 {
		s2[v] = struct{}{}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var all = len(s1)
	for k := range s2 {
		if _, ok := s1[k]; !ok {
			all++
		}
	}
	return min(all, min(len(s1), half)+min(len(s2), half))
}
