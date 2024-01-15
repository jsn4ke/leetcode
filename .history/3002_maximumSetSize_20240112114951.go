package leatcode

func maximumSetSize(nums1 []int, nums2 []int) int {
	var (
		s1 = map[int]struct{}{}
		s2 = map[int]struct{}{}
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
	var all int
	var s1n2 int
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			s1n2++
			all++
		}
	}
	s1n2 = min(s1n2, len(nums1))
	var s2n1 int
	for k := range s2 {
		if _, ok := s1[k]; !ok {
			s2n1++
		}
	}
	s2n1 = min(s2n1, len(nums1))

	return 0
}
