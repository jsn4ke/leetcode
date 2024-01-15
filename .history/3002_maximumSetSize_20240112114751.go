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
	var s1n2 int
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			s1n2++
		}
	}
	return 0
}
